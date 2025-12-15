package image

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/errorz"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/models"
)

func (s *service) GetProcessed(ctx context.Context, id uuid.UUID) (io.ReadCloser, string, models.Status, int64, error) {
	imageInfo, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, "", "", 0, err
	}
	if imageInfo.Status != models.StatusDone {
		return nil, "", imageInfo.Status, 0, errorz.ErrImageIsNotDone
	}

	img, contentType, size, err := s.storage.Get(ctx, fmt.Sprintf("%s_processed", id.String()))
	if err != nil {
		return nil, "", "", 0, err
	}

	return img, contentType, imageInfo.Status, size, nil
}
