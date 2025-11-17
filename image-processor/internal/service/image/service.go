package imageuploader

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/models"
)

type imageStorage interface {
	Upload(ctx context.Context, file io.Reader, size int64, filename, contentType string) error
}

type kafkaProducer interface {
	SendImage(ctx context.Context, id uuid.UUID) error
}

type imageRepo interface {
	Create(ctx context.Context, filename string) (*models.Image, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, newStatus string) error
}

type service struct {
	storage       imageStorage
	repo          imageRepo
	kafkaProducer kafkaProducer
}

func New(storage imageStorage, repo imageRepo, kafkaProducer kafkaProducer) *service {
	return &service{
		storage:       storage,
		repo:          repo,
		kafkaProducer: kafkaProducer,
	}
}
