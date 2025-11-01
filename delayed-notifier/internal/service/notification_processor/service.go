package notification_processor

import (
	"context"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
)

type repo interface {
	Get(ctx context.Context, id uuid.UUID) (*models.Notification, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.NotificationStatus) error
}

type telegramSender interface {
	SendNotification(ctx context.Context, title, content, receiver string) error
}

type Service struct {
	telegramSender telegramSender
	repo           repo
}

func New(
	telegramSender telegramSender,
	repo repo,
) *Service {
	return &Service{
		telegramSender: telegramSender,
		repo:           repo,
	}
}
