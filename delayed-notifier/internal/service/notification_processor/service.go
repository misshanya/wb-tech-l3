package notification_processor

import (
	"context"
)

type telegramSender interface {
	SendNotification(ctx context.Context, title, content, receiver string) error
}

type Service struct {
	telegramSender telegramSender
}

func New(telegramSender telegramSender) *Service {
	return &Service{
		telegramSender: telegramSender,
	}
}
