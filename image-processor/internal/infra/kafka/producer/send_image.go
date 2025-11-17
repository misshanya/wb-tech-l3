package producer

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (p *producer) SendImage(ctx context.Context, id uuid.UUID) error {
	idBytes, err := id.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal id: %w", err)
	}

	err = p.p.SendWithRetry(ctx, p.retry, nil, idBytes)
	if err != nil {
		return fmt.Errorf("failed to send image id to kafka: %w", err)
	}

	return nil
}
