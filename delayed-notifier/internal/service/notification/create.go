package notification

import (
	"context"
	"fmt"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
)

func (s *Service) Create(ctx context.Context, n *models.Notification) (*models.Notification, error) {
	notification, err := s.repo.Create(ctx, n)
	if err != nil {
		return nil, fmt.Errorf("failed to create notification: %w", err)
	}

	err = s.producer.SendNotification(notification)
	if err != nil {
		return nil, fmt.Errorf("failed to send notification: %w", err)
	}

	return notification, nil
}
