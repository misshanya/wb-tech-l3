package notification

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/events"
	"github.com/wb-go/wbf/zlog"
)

func (c *Consumer) ConsumeMessages() {
	ch := make(chan []byte, c.workers)

	// Goroutine to consume
	go func() {
		if err := c.consumer.ConsumeWithRetry(ch, c.retry); err != nil {
			zlog.Logger.Warn().
				Err(err).
				Msg("failed to consume messages")
		}
		close(ch)
	}()

	// Run consumer workers
	wg := &sync.WaitGroup{}
	for range c.workers {
		wg.Go(func() {
			for msg := range ch {
				zlog.Logger.Info().
					Msg("received message")

				var notification events.Notification
				if err := json.Unmarshal(msg, &notification); err != nil {
					zlog.Logger.Warn().
						Err(err).
						Msg("failed to unmarshal notification")
				}

				procCtx, cancel := context.WithTimeout(context.Background(), c.processTimeout)

				if err := c.processor.ProcessNotification(procCtx, &notification); err != nil {
					zlog.Logger.Error().
						Err(err).
						Msg("failed to process notification")
				} else {
					zlog.Logger.Info().
						Any("notification", notification).
						Msg("processed notification")
				}

				cancel()
			}
		})
	}

	wg.Wait()
}
