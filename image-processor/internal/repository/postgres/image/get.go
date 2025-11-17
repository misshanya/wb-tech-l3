package image

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/models"
)

func (r *repo) Get(ctx context.Context, id uuid.UUID) (*models.Image, error) {
	image, err := r.queries.GetImage(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get image from db: %w", err)
	}

	return &models.Image{
		ID:               image.ID,
		OriginalFilename: image.OriginalFilename,
		Status:           models.Status(image.Status),
	}, nil
}
