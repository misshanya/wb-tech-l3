package app

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/config"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/db/ent"
	producer "github.com/misshanya/wb-tech-l3/delayed-notifier/internal/infra/rabbitmq/producer/notification"
	telegramsender "github.com/misshanya/wb-tech-l3/delayed-notifier/internal/infra/telegram/notification"
	notificationrepo "github.com/misshanya/wb-tech-l3/delayed-notifier/internal/repository/notification"
	notificationservice "github.com/misshanya/wb-tech-l3/delayed-notifier/internal/service/notification"
	notificationprocessor "github.com/misshanya/wb-tech-l3/delayed-notifier/internal/service/notification_processor"
	handler "github.com/misshanya/wb-tech-l3/delayed-notifier/internal/transport/http/v1/notification"
	consumer "github.com/misshanya/wb-tech-l3/delayed-notifier/internal/transport/rabbitmq/consumer/notification"
	"github.com/wb-go/wbf/dbpg"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"

	"net/http"

	"github.com/wb-go/wbf/rabbitmq"
)

type App struct {
	cfg              *config.Config
	rabbitMQConn     *rabbitmq.Connection
	rabbitMQProducer *producer.Producer
	rabbitMQConsumer *consumer.Consumer
	ginextEngine     *ginext.Engine
	httpSrv          *http.Server
	telegramSender   *telegramsender.Sender
	pgConn           *dbpg.DB
	entClient        *ent.Client
}

// New creates and initializes a new instance of App
func New(ctx context.Context, cfg *config.Config) (*App, error) {
	a := &App{
		cfg: cfg,
	}

	if err := a.initRabbitMQ(); err != nil {
		return nil, fmt.Errorf("failed to init rabbitmq: %w", err)
	}

	if err := a.initRabbitMQProducer(); err != nil {
		return nil, fmt.Errorf("failed to init rabbitmq producer: %w", err)
	}

	a.initTelegramSender()

	if err := a.initDB(); err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	if err := a.migrateDB(ctx); err != nil {
		return nil, fmt.Errorf("failed to migrate db: %w", err)
	}

	repo := notificationrepo.New(a.entClient)

	notificationProc := notificationprocessor.New(a.telegramSender, repo)
	if err := a.initRabbitMQConsumer(notificationProc); err != nil {
		return nil, fmt.Errorf("failed to init rabbitmq consumer: %w", err)
	}
	svc := notificationservice.New(a.rabbitMQProducer, repo)

	h := handler.New(svc)

	a.initGinext()

	notifyGroup := a.ginextEngine.Group("/api/v1/notify")

	h.Setup(notifyGroup)

	return a, nil
}

// Start performs a start of all functional services
func (a *App) Start(errChan chan<- error) {
	zlog.Logger.Info().Msg("starting...")
	go a.rabbitMQConsumer.ConsumeMessages()
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

	zlog.Logger.Info().Msg("Closing RabbitMQ connection...")
	if err := a.rabbitMQConn.Close(); err != nil {
		stopErr = errors.Join(stopErr, fmt.Errorf("failed to close RabbitMQ connection: %w", err))
	}

	if stopErr != nil {
		return stopErr
	}

	zlog.Logger.Info().Msg("Stopped gracefully!")
	return nil
}

func (a *App) initRabbitMQ() error {
	conn, err := rabbitmq.Connect(a.cfg.RabbitMQ.URL, a.cfg.RabbitMQ.ConnectRetries, a.cfg.RabbitMQ.ConnectRetryPause)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	a.rabbitMQConn = conn
	return nil
}

func (a *App) initRabbitMQProducer() error {
	channel, err := a.rabbitMQConn.Channel()
	if err != nil {
		return fmt.Errorf("failed to create RabbitMQ channel: %w", err)
	}

	p, err := producer.New(
		channel,
		a.cfg.RabbitMQ.ExchangeName,
		a.cfg.RabbitMQ.RoutingKey,
		a.cfg.RabbitMQ.Producer.Retry.Attempts,
		a.cfg.RabbitMQ.Producer.Retry.Delay,
		a.cfg.RabbitMQ.Producer.Retry.Backoff,
	)
	if err != nil {
		return fmt.Errorf("failed to create rabbitmq producer: %w", err)
	}

	a.rabbitMQProducer = p

	return nil
}

func (a *App) initRabbitMQConsumer(processor *notificationprocessor.Service) error {
	channel, err := a.rabbitMQConn.Channel()
	if err != nil {
		return fmt.Errorf("failed to create RabbitMQ channel: %w", err)
	}

	c, err := consumer.New(
		channel,
		a.cfg.RabbitMQ.ExchangeName,
		a.cfg.RabbitMQ.Consumer.QueueName,
		a.cfg.RabbitMQ.RoutingKey,
		a.cfg.RabbitMQ.Consumer.Retry.Attempts,
		a.cfg.RabbitMQ.Consumer.Retry.Delay,
		a.cfg.RabbitMQ.Consumer.Retry.Backoff,
		a.cfg.RabbitMQ.Consumer.Workers,
		processor,
		a.cfg.RabbitMQ.Consumer.ProcessMessageTimeout,
	)
	if err != nil {
		return fmt.Errorf("failed to create rabbitmq consumer: %w", err)
	}

	a.rabbitMQConsumer = c

	return nil
}

func (a *App) initTelegramSender() {
	httpClient := &http.Client{}
	s := telegramsender.New(
		httpClient,
		a.cfg.TelegramSender.BotApiToken,
		a.cfg.TelegramSender.Retry.Attempts,
		a.cfg.TelegramSender.Retry.Delay,
		a.cfg.TelegramSender.Retry.Backoff,
	)
	a.telegramSender = s
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
