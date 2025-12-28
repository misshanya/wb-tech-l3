package producer

import (
	"sync"
	"time"

	"github.com/wb-go/wbf/kafka"
	"github.com/wb-go/wbf/retry"
)

type Producer struct {
	p           *kafka.Producer
	retry       retry.Strategy
	workersChan chan []byte
	wg          *sync.WaitGroup
}

func New(
	p *kafka.Producer,
	retryAttempts int,
	retryDelay time.Duration,
	retryBackoff float64,
	bufferSize int,
	numWorkers int,
) *Producer {
	pr := &Producer{
		p: p,
		retry: retry.Strategy{
			Attempts: retryAttempts,
			Delay:    retryDelay,
			Backoff:  retryBackoff,
		},
		workersChan: make(chan []byte, bufferSize),
		wg:          &sync.WaitGroup{},
	}

	for i := range numWorkers {
		pr.wg.Go(func() {
			pr.worker(i)
		})
	}

	return pr
}

func (p *Producer) Stop() {
	close(p.workersChan)
	p.wg.Wait()
}
