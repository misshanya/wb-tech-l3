package notification

import (
	"fmt"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
)

func (s *Service) Create(n *models.Notification) (*models.Notification, error) {
	// TODO: create notification in database
	// Now it just sends notification to the message broker
	err := s.producer.SendNotification(n)
	if err != nil {
		return nil, fmt.Errorf("failed to send notification: %w", err)
	}

	return n, nil
}
