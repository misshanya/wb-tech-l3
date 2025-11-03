package comment

import (
	"context"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
)

type repo interface {
	Create(ctx context.Context, c *models.Comment) (*models.Comment, error)
	GetDerivatives(ctx context.Context, id uuid.UUID) ([]*models.Comment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo repo
}

func New(repo repo) *service {
	return &service{repo: repo}
}
