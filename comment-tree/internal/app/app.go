package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/config"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/db/ent"
	commentrepo "github.com/misshanya/wb-tech-l3/comment-tree/internal/repository/comment"
	commentservice "github.com/misshanya/wb-tech-l3/comment-tree/internal/service/comment"
	commenthandler "github.com/misshanya/wb-tech-l3/comment-tree/internal/transport/http/v1/comment"
	"github.com/wb-go/wbf/dbpg"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

type App struct {
	cfg          *config.Config
	ginextEngine *ginext.Engine
	httpSrv      *http.Server
	pgConn       *dbpg.DB
	entClient    *ent.Client
}

// New creates and initializes a new instance of App
func New(ctx context.Context, cfg *config.Config) (*App, error) {
	a := &App{
		cfg: cfg,
	}

	if err := a.initDB(); err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	if err := a.migrateDB(ctx); err != nil {
		return nil, fmt.Errorf("failed to migrate db: %w", err)
	}

	commentRepo := commentrepo.New(a.entClient)
	commentService := commentservice.New(commentRepo)
	commentHandler := commenthandler.New(commentService)

	a.initGinext()

	commentGroup := a.ginextEngine.Group("/api/v1/comment")

	commentHandler.Setup(commentGroup)

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
	if err := a.entClient.Close(); err != nil {
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

	drv := entsql.OpenDB(dialect.Postgres, db.Master)
	a.entClient = ent.NewClient(ent.Driver(drv))

	return nil
}

func (a *App) migrateDB(ctx context.Context) error {
	return a.entClient.Schema.Create(ctx)
}

func (a *App) initGinext() {
	a.ginextEngine = ginext.New(gin.ReleaseMode)
	a.httpSrv = &http.Server{
		Addr:    a.cfg.Server.Addr,
		Handler: a.ginextEngine,
	}
}
