package mappers

import (
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/db/ent"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
)

func EntNotificationToModel(e *ent.Notification) *models.Notification {
	return &models.Notification{
		ID:          e.ID,
		ScheduledAt: e.ScheduledAt,
		Title:       e.Title,
		Content:     e.Content,
		Channel:     e.Channel,
		Receiver:    e.Receiver,
		Status:      models.NotificationStatus(e.Status),
	}
}
