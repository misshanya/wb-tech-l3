package comment

import (
	"context"

	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
)

func (s *service) Create(ctx context.Context, c *models.Comment) (*models.Comment, error) {
	return s.repo.Create(ctx, c)
}
