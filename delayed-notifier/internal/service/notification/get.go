package notification

import (
	"context"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
)

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*models.Notification, error) {
	return s.repo.Get(ctx, id)
}
