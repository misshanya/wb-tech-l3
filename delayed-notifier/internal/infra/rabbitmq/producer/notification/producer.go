package notification

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wb-go/wbf/rabbitmq"
	"github.com/wb-go/wbf/retry"
)

type Producer struct {
	channel    *rabbitmq.Channel
	publisher  *rabbitmq.Publisher
	routingKey string
	retry      retry.Strategy
}

func New(
	channel *rabbitmq.Channel,
	exchangeName, routingKey string,
	retryAttempts int,
	retryDelay time.Duration,
	retryBackoff float64,
) (*Producer, error) {
	p := &Producer{
		channel:    channel,
		routingKey: routingKey,
		retry: retry.Strategy{
			Attempts: retryAttempts,
			Delay:    retryDelay,
			Backoff:  retryBackoff,
		},
	}

	exchange := rabbitmq.NewExchange(exchangeName, "x-delayed-message")
	exchange.Durable = true
	exchange.Args = amqp.Table{
		"x-delayed-type": "direct",
	}

	err := exchange.BindToChannel(p.channel)
	if err != nil {
		return nil, fmt.Errorf("failed to declare an exchange: %s", err)
	}

	p.publisher = rabbitmq.NewPublisher(p.channel, exchangeName)

	return p, nil
}

func (p *Producer) Close() error {
	return p.channel.Close()
}
