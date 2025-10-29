package config

import (
	"time"

	"github.com/wb-go/wbf/config"
)

type Config struct {
	Server         server
	RabbitMQ       rabbitmq
	TelegramSender telegramSender
}

type server struct {
	Addr string
}

type sender struct {
	Retry retry
}

type telegramSender struct {
	sender
	BotApiToken string
}

type rabbitmq struct {
	URL               string
	ConnectRetries    int
	ConnectRetryPause time.Duration
	ExchangeName      string
	RoutingKey        string
	Producer          rabbitMQProducer
	Consumer          rabbitMQConsumer
}

type rabbitMQProducer struct {
	Retry retry
}

type rabbitMQConsumer struct {
	Retry                 retry
	Workers               int
	ProcessMessageTimeout time.Duration
	QueueName             string
}

type retry struct {
	Attempts int
	Delay    time.Duration
	Backoff  float64
}

func New() *Config {
	c := config.New()
	c.EnableEnv("")
	_ = c.LoadEnvFiles(".env", ".env.example")

	cfg := &Config{
		Server: server{
			Addr: c.GetString("server.addr"),
		},
		RabbitMQ: rabbitmq{
			URL:               c.GetString("rabbitmq.url"),
			ConnectRetries:    c.GetInt("rabbitmq.connect_retries"),
			ConnectRetryPause: c.GetDuration("rabbitmq.connect_retry_pause"),
			ExchangeName:      c.GetString("rabbitmq.exchange_name"),
			RoutingKey:        c.GetString("rabbitmq.routing_key"),
			Producer: rabbitMQProducer{
				Retry: retry{
					Attempts: c.GetInt("rabbitmq.producer.retry.attempts"),
					Delay:    c.GetDuration("rabbitmq.producer.retry.delay"),
					Backoff:  c.GetFloat64("rabbitmq.producer.retry.backoff"),
				},
			},
			Consumer: rabbitMQConsumer{
				Retry: retry{
					Attempts: c.GetInt("rabbitmq.consumer.retry.attempts"),
					Delay:    c.GetDuration("rabbitmq.consumer.retry.delay"),
					Backoff:  c.GetFloat64("rabbitmq.consumer.retry.backoff"),
				},
				Workers:               c.GetInt("rabbitmq.consumer.workers"),
				ProcessMessageTimeout: c.GetDuration("rabbitmq.consumer.process_message_timeout"),
				QueueName:             c.GetString("rabbitmq.consumer.queue_name"),
			},
		},
		TelegramSender: telegramSender{
			sender: sender{
				Retry: retry{
					Attempts: c.GetInt("telegram.sender.retry.attempts"),
					Delay:    c.GetDuration("telegram.sender.retry.delay"),
					Backoff:  c.GetFloat64("telegram.sender.retry.backoff"),
				},
			},
			BotApiToken: c.GetString("telegram.sender.bot_api_token"),
		},
	}

	return cfg
}
