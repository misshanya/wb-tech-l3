package notification

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/events"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wb-go/wbf/rabbitmq"
)

func (p *Producer) SendNotification(notification *models.Notification) error {
	notificationDTO := &events.Notification{
		ID:       notification.ID,
		Title:    notification.Title,
		Content:  notification.Content,
		Channel:  notification.Channel,
		Receiver: notification.Receiver,
	}
	body, err := json.Marshal(notificationDTO)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	delay := notification.ScheduledAt.Sub(time.Now()).Milliseconds()
	headers := amqp.Table{
		"x-delay": delay,
	}

	err = p.publisher.PublishWithRetry(
		body,
		p.routingKey,
		"application/json",
		p.retry,
		rabbitmq.PublishingOptions{
			Headers: headers,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	return nil
}
