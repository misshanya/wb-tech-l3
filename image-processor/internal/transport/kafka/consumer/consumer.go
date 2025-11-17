package consumer

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wb-go/wbf/kafka"
	"github.com/wb-go/wbf/retry"
)

type service interface {
	ProcessImage(ctx context.Context, id uuid.UUID) error
}

type Consumer struct {
	service             service
	kafkaConsumer       *kafka.Consumer
	processImageTimeout time.Duration
	retry               retry.Strategy
}

func New(
	service service,
	kafkaConsumer *kafka.Consumer,
	processImageTimeout time.Duration,
	retryAttempts int,
	retryDelay time.Duration,
	retryBackoff float64,
) *Consumer {
	return &Consumer{
		service:             service,
		kafkaConsumer:       kafkaConsumer,
		processImageTimeout: processImageTimeout,
		retry: retry.Strategy{
			Attempts: retryAttempts,
			Delay:    retryDelay,
			Backoff:  retryBackoff,
		},
	}
}
