package producer

import (
	"context"
	"time"

	"github.com/wb-go/wbf/zlog"
)

func (p *Producer) worker(id int) {
	for msg := range p.workersChan {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			err := p.p.SendWithRetry(ctx, p.retry, nil, msg)
			if err != nil {
				zlog.Logger.Warn().Int("worker_id", id).Err(err).Msg("failed to send message to Kafka")
			}
		}()
	}
}
