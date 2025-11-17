package producer

import (
	"time"

	"github.com/wb-go/wbf/kafka"
	"github.com/wb-go/wbf/retry"
)

type producer struct {
	p     *kafka.Producer
	retry retry.Strategy
}

func New(
	p *kafka.Producer,
	retryAttempts int,
	retryDelay time.Duration,
	retryBackoff float64,
) *producer {
	return &producer{
		p: p,
		retry: retry.Strategy{
			Attempts: retryAttempts,
			Delay:    retryDelay,
			Backoff:  retryBackoff,
		},
	}
}
