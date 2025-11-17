package image

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/db/sqlc/storage"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/errorz"
)

func (r *repo) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	err := r.queries.UpdateStatus(ctx, storage.UpdateStatusParams{
		ID:     id,
		Status: status,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorz.ErrImageNotFound
		}
		return fmt.Errorf("failed to update image status in db: %w", err)
	}
	return nil
}
