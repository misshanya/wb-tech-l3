package notification

import (
	"context"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
)

func (s *Service) Cancel(ctx context.Context, id uuid.UUID) error {
	return s.repo.UpdateStatus(ctx, id, models.StatusCancelled)
}
