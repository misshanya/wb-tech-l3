package consumer

import (
	"context"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/wb-go/wbf/zlog"
)

func (c *Consumer) Consume(ctx context.Context) {
	msgCh := make(chan kafka.Message)
	c.kafkaConsumer.StartConsuming(ctx, msgCh, c.retry)

	for msg := range msgCh {
		func() {
			zlog.Logger.Debug().Str("msg", string(msg.Value)).Msg("message received")
			id, err := uuid.FromBytes(msg.Value)
			if err != nil {
				zlog.Logger.Warn().Err(err).Msg("invalid id")
				return
			}

			ctx, cancel := context.WithTimeout(ctx, c.processImageTimeout)
			defer cancel()
			err = c.service.ProcessImage(ctx, id)
			if err != nil {
				zlog.Logger.Warn().Err(err).Msg("process image error")
				return
			}

			err = c.kafkaConsumer.Commit(ctx, msg)
			if err != nil {
				zlog.Logger.Warn().Err(err).Msg("commit error")
				return
			}
		}()
	}
}
