package config

import (
	"time"

	"github.com/wb-go/wbf/config"
)

type Config struct {
	Server   server
	Postgres postgres
	MinIO    minio
	Kafka    kafka
}

type server struct {
	Addr string
}

type postgres struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type minio struct {
	Endpoint   string
	AccessKey  string
	SecretKey  string
	BucketName string
}

type retry struct {
	Attempts int
	Delay    time.Duration
	Backoff  float64
}

type kafka struct {
	Addr     string
	Topic    string
	Producer kafkaProducer
}

type kafkaProducer struct {
	Retry retry
}

func New() *Config {
	c := config.New()
	c.EnableEnv("")
	_ = c.LoadEnvFiles(".env", ".env.example")

	cfg := &Config{
		Server: server{
			Addr: c.GetString("server.addr"),
		},
		Postgres: postgres{
			URL:             c.GetString("postgres.url"),
			MaxOpenConns:    c.GetInt("postgres.max_open_conns"),
			MaxIdleConns:    c.GetInt("postgres.max_idle_conns"),
			ConnMaxLifetime: c.GetDuration("postgres.conn_max_lifetime"),
		},
		MinIO: minio{
			Endpoint:   c.GetString("minio.endpoint"),
			AccessKey:  c.GetString("minio.access_key"),
			SecretKey:  c.GetString("minio.secret_key"),
			BucketName: c.GetString("minio.bucket_name"),
		},
		Kafka: kafka{
			Addr:  c.GetString("kafka.addr"),
			Topic: c.GetString("kafka.topic"),
			Producer: kafkaProducer{
				Retry: retry{
					Attempts: c.GetInt("kafka.producer.retry.attempts"),
					Delay:    c.GetDuration("kafka.producer.retry.delay"),
					Backoff:  c.GetFloat64("kafka.producer.retry.backoff"),
				},
			},
		},
	}

	return cfg
}
