package producer

import (
	"fmt"

	"github.com/google/uuid"
)

func (p *Producer) SendImage(id uuid.UUID) error {
	idBytes, err := id.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal id: %w", err)
	}

	p.workersChan <- idBytes

	return nil
}
