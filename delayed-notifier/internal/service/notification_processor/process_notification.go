package notification_processor

import (
	"context"
	"fmt"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/errorz"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/events"
)

func (s *Service) ProcessNotification(ctx context.Context, notification *events.Notification) error {
	var err error

	switch notification.Channel {
	case "telegram":
		err = s.telegramSender.SendNotification(ctx, notification.Title, notification.Content, notification.Receiver)
	default:
		return errorz.ChannelNotFoundError
	}
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	// TODO: write to db that notification is sent now

	return nil
}
