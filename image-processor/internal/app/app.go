package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/config"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/db"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/db/sqlc/storage"
	"github.com/wb-go/wbf/dbpg"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/kafka"
	"github.com/wb-go/wbf/zlog"

	minioimagerepo "github.com/misshanya/wb-tech-l3/image-processor/internal/repository/minio/image"
	pgimagerepo "github.com/misshanya/wb-tech-l3/image-processor/internal/repository/postgres/image"
	imageservice "github.com/misshanya/wb-tech-l3/image-processor/internal/service/image"
	imagehandler "github.com/misshanya/wb-tech-l3/image-processor/internal/transport/http/v1/image"

	kafkaproducer "github.com/misshanya/wb-tech-l3/image-processor/internal/infra/kafka/producer"
)

type App struct {
	cfg           *config.Config
	ginextEngine  *ginext.Engine
	httpSrv       *http.Server
	pgConn        *dbpg.DB
	minioClient   *minio.Client
	kafkaProducer *kafka.Producer
}

// New creates and initializes a new instance of App
func New(ctx context.Context, cfg *config.Config) (*App, error) {
	a := &App{
		cfg: cfg,
	}

	if err := a.initDB(); err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	if err := a.migrateDB(); err != nil {
		return nil, fmt.Errorf("failed to migrate db: %w", err)
	}

	if err := a.initMinIO(ctx); err != nil {
		return nil, fmt.Errorf("failed to init minio: %w", err)
	}

	a.initKafka()

	queries := storage.New(a.pgConn.Master)

	minioImageRepo := minioimagerepo.New(a.minioClient, a.cfg.MinIO.BucketName)
	pgImageRepo := pgimagerepo.New(queries)

	kafkaProducer := kafkaproducer.New(
		a.kafkaProducer,
		a.cfg.Kafka.Producer.Retry.Attempts,
		a.cfg.Kafka.Producer.Retry.Delay,
		a.cfg.Kafka.Producer.Retry.Backoff,
	)

	imageService := imageservice.New(minioImageRepo, pgImageRepo, kafkaProducer)

	imageHandler := imagehandler.New(imageService)

	a.initGinext()

	apiGroup := a.ginextEngine.Group("/api/v1")

	imageHandler.Setup(apiGroup)

	return a, nil
}

// Start performs a start of all functional services
func (a *App) Start(errChan chan<- error) {
	zlog.Logger.Info().Msg("starting...")
	if err := a.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		errChan <- err
	}
}

func (a *App) Stop() error {
	zlog.Logger.Info().Msg("[!] Shutting down...")

	var stopErr error

	zlog.Logger.Info().Msg("Stopping http server...")
	if err := a.httpSrv.Close(); err != nil {
		stopErr = errors.Join(stopErr, fmt.Errorf("failed to stop http server: %w", err))
	}

	zlog.Logger.Info().Msg("Closing db connection...")
	if err := a.pgConn.Master.Close(); err != nil {
		stopErr = errors.Join(stopErr, fmt.Errorf("failed to close db connection: %w", err))
	}

	if stopErr != nil {
		return stopErr
	}

	zlog.Logger.Info().Msg("Stopped gracefully!")
	return nil
}

func (a *App) initDB() error {
	db, err := dbpg.New(a.cfg.Postgres.URL, nil, &dbpg.Options{
		MaxOpenConns:    a.cfg.Postgres.MaxOpenConns,
		MaxIdleConns:    a.cfg.Postgres.MaxIdleConns,
		ConnMaxLifetime: a.cfg.Postgres.ConnMaxLifetime,
	})
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	a.pgConn = db

	return nil
}

func (a *App) migrateDB() error {
	return db.Migrate(a.pgConn.Master)
}

func (a *App) initMinIO(ctx context.Context) error {
	client, err := minio.New(a.cfg.MinIO.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(a.cfg.MinIO.AccessKey, a.cfg.MinIO.SecretKey, ""),
	})
	if err != nil {
		return fmt.Errorf("failed to create minio client: %w", err)
	}
	a.minioClient = client
	exists, err := client.BucketExists(ctx, a.cfg.MinIO.BucketName)
	if err != nil {
		return fmt.Errorf("failed to check if bucket exists: %w", err)
	}
	if !exists {
		if err := client.MakeBucket(ctx, a.cfg.MinIO.BucketName, minio.MakeBucketOptions{}); err != nil {
			return fmt.Errorf("failed to create bucket: %w", err)
		}
	}
	return nil
}

func (a *App) initKafka() {
	a.kafkaProducer = kafka.NewProducer([]string{a.cfg.Kafka.Addr}, a.cfg.Kafka.Topic)
	a.kafkaProducer.Writer.AllowAutoTopicCreation = true
}

func (a *App) initGinext() {
	a.ginextEngine = ginext.New(gin.ReleaseMode)
	a.httpSrv = &http.Server{
		Addr:    a.cfg.Server.Addr,
		Handler: a.ginextEngine,
	}
}
