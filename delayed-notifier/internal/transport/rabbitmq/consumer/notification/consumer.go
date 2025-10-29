package notification

import (
	"context"
	"fmt"
	"time"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/events"
	"github.com/wb-go/wbf/rabbitmq"
	"github.com/wb-go/wbf/retry"
)

type notificationsProcessor interface {
	ProcessNotification(ctx context.Context, n *events.Notification) error
}

type Consumer struct {
	channel        *rabbitmq.Channel
	consumer       *rabbitmq.Consumer
	routingKey     string
	retry          retry.Strategy
	workers        int
	processor      notificationsProcessor
	processTimeout time.Duration
}

func New(
	channel *rabbitmq.Channel,
	exchangeName string,
	queueName string,
	routingKey string,
	retryAttempts int,
	retryDelay time.Duration,
	retryBackoff float64,
	workers int,
	processor notificationsProcessor,
	processTimeout time.Duration,
) (*Consumer, error) {
	c := &Consumer{
		channel:    channel,
		routingKey: routingKey,
		retry: retry.Strategy{
			Attempts: retryAttempts,
			Delay:    retryDelay,
			Backoff:  retryBackoff,
		},
		workers:        workers,
		processor:      processor,
		processTimeout: processTimeout,
	}

	queueManager := rabbitmq.NewQueueManager(c.channel)
	queue, err := queueManager.DeclareQueue(
		queueName,
		rabbitmq.QueueConfig{
			Durable: true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	err = c.channel.QueueBind(
		queue.Name,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	c.consumer = rabbitmq.NewConsumer(c.channel, &rabbitmq.ConsumerConfig{
		Queue:    queueName,
		Consumer: "notifications",
	})

	return c, nil
}

func (c *Consumer) Close() error {
	return c.channel.Close()
}
