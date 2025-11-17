package image

import (
	"context"
	"fmt"

	"github.com/misshanya/wb-tech-l3/image-processor/internal/models"
)

func (r *repo) Create(ctx context.Context, filename string) (*models.Image, error) {
	image, err := r.queries.CreateImage(ctx, filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create image in db: %w", err)
	}
	return &models.Image{
		ID:               image.ID,
		OriginalFilename: image.OriginalFilename,
		Status:           models.Status(image.Status),
	}, nil
}
