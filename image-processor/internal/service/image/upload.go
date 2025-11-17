package imageuploader

import (
	"context"
	"io"

	"github.com/misshanya/wb-tech-l3/image-processor/internal/models"
)

func (s *service) Upload(ctx context.Context, content io.Reader, size int64, filename, contentType string) (*models.Image, error) {
	imageInfo, err := s.repo.Create(ctx, filename)
	if err != nil {
		return nil, err
	}

	err = s.storage.Upload(ctx, content, size, imageInfo.ID.String(), contentType)
	if err != nil {
		return nil, err
	}

	err = s.kafkaProducer.SendImage(ctx, imageInfo.ID)
	if err != nil {
		return nil, err
	}

	imageInfo.Status = models.StatusPending
	err = s.repo.UpdateStatus(ctx, imageInfo.ID, string(imageInfo.Status))
	if err != nil {
		return nil, err
	}

	return imageInfo, nil
}
