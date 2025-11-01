package notification

import (
	"context"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/db/ent"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/errorz"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
)

func (r *repo) UpdateStatus(ctx context.Context, id uuid.UUID, status models.NotificationStatus) error {
	_, err := r.client.Notification.
		UpdateOneID(id).
		SetStatus(string(status)).
		Save(ctx)
	if ent.IsNotFound(err) {
		return errorz.NotificationNotFound
	}
	return err
}
