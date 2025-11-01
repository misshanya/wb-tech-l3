package notification

import (
	"context"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/repository/mappers"
)

func (r *repo) Create(ctx context.Context, n *models.Notification) (*models.Notification, error) {
	created, err := r.client.Notification.
		Create().
		SetScheduledAt(n.ScheduledAt).
		SetTitle(n.Title).
		SetContent(n.Content).
		SetChannel(n.Channel).
		SetReceiver(n.Receiver).
		SetStatus(string(models.StatusScheduled)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mappers.EntNotificationToModel(created), nil
}
