package notification

import (
	"context"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/db/ent"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/db/ent/notification"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/errorz"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/repository/mappers"
)

func (r *repo) Get(ctx context.Context, id uuid.UUID) (*models.Notification, error) {
	n, err := r.client.Notification.
		Query().
		Where(notification.ID(id)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.NotificationNotFound
		}
		return nil, err
	}

	return mappers.EntNotificationToModel(n), nil
}
