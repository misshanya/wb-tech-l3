package comment

import (
	"context"

	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
)

func (s *service) Search(ctx context.Context, query string, limit, offset int32) ([]*models.Comment, error) {
	return s.repo.Search(ctx, query, limit, offset)
}
