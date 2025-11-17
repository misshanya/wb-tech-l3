package imageprocessor

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/models"
)

type imageStorage interface {
	Upload(ctx context.Context, file io.Reader, size int64, filename, contentType string) error
	Get(ctx context.Context, filename string) (io.ReadCloser, string, int64, error)
}

type repo interface {
	Get(ctx context.Context, id uuid.UUID) (*models.Image, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, newStatus string) error
}

type service struct {
	repo         repo
	imageStorage imageStorage
	resizeFactor int
}

func New(
	repo repo,
	imageStorage imageStorage,
	resizeFactor int,
) *service {
	return &service{
		repo:         repo,
		imageStorage: imageStorage,
		resizeFactor: resizeFactor,
	}
}
