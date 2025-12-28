package config

import (
	"time"

	"github.com/wb-go/wbf/config"
)

type Config struct {
	Server          server
	Postgres        postgres
	MinIO           minio
	Kafka           kafka
	ImageProcessing imageProcessing
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
	Consumer kafkaConsumer
}

type kafkaProducer struct {
	Retry      retry
	BufferSize int
	NumWorkers int
}

type kafkaConsumer struct {
	Retry   retry
	GroupID string
}

type imageProcessing struct {
	Timeout       time.Duration
	ResizeFactor  int
	WatermarkPath string
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
				BufferSize: c.GetInt("kafka.producer.buffer.size"),
				NumWorkers: c.GetInt("kafka.producer.num_workers"),
			},
			Consumer: kafkaConsumer{
				Retry: retry{
					Attempts: c.GetInt("kafka.consumer.retry.attempts"),
					Delay:    c.GetDuration("kafka.consumer.retry.delay"),
					Backoff:  c.GetFloat64("kafka.consumer.retry.backoff"),
				},
				GroupID: c.GetString("kafka.consumer.group_id"),
			},
		},
		ImageProcessing: imageProcessing{
			Timeout:       c.GetDuration("image_processing.timeout"),
			ResizeFactor:  c.GetInt("image_processing.resize_factor"),
			WatermarkPath: c.GetString("image_processing.watermark_path"),
		},
	}

	return cfg
}
