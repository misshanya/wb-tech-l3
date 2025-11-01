package notification

import (
	"context"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
)

type repo interface {
	Create(ctx context.Context, n *models.Notification) (*models.Notification, error)
	Get(ctx context.Context, id uuid.UUID) (*models.Notification, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.NotificationStatus) error
}

type producer interface {
	SendNotification(notification *models.Notification) error
}

type Service struct {
	producer producer
	repo     repo
}

func New(
	producer producer,
	repo repo,
) *Service {
	return &Service{
		producer,
		repo,
	}
}
