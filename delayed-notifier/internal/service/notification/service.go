package notification

import "github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"

type producer interface {
	SendNotification(notification *models.Notification) error
}

type Service struct {
	producer producer
}

func New(producer producer) *Service {
	return &Service{producer}
}
