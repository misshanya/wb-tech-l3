package image

import (
	"context"
	"fmt"

	"github.com/misshanya/wb-tech-l3/image-processor/internal/models"
)

func (r *repo) Create(ctx context.Context) (*models.Image, error) {
	image, err := r.queries.CreateImage(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create image in db: %w", err)
	}
	return &models.Image{
		ID:     image.ID,
		Status: models.Status(image.Status),
	}, nil
}
