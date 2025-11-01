package notification_processor

import (
	"context"
	"fmt"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/errorz"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/events"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
)

func (s *Service) ProcessNotification(ctx context.Context, notification *events.Notification) error {
	n, err := s.repo.Get(ctx, notification.ID)
	if err != nil {
		return fmt.Errorf("failed to get notification: %w", err)
	}

	if n.Status != models.StatusScheduled {
		return nil
	}

	err = s.repo.UpdateStatus(ctx, notification.ID, models.StatusProcessing)
	if err != nil {
		return fmt.Errorf("failed to update status to processing: %w", err)
	}

	switch notification.Channel {
	case "telegram":
		err = s.telegramSender.SendNotification(ctx, notification.Title, notification.Content, notification.Receiver)
	default:
		return errorz.ChannelNotFoundError
	}
	if err != nil {
		err = s.repo.UpdateStatus(ctx, notification.ID, models.StatusFailed)
		if err != nil {
			return fmt.Errorf("failed to update status to failed: %w", err)
		}
		return fmt.Errorf("failed to send notification: %w", err)
	}

	err = s.repo.UpdateStatus(ctx, notification.ID, models.StatusDelivered)
	if err != nil {
		return fmt.Errorf("failed to update status to delivered: %w", err)
	}

	return nil
}
