package comment

import (
	"context"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
)

func (s *service) Get(ctx context.Context, id uuid.UUID, limit, offset int32) ([]*models.Comment, error) {
	return s.repo.GetDerivatives(ctx, id, limit, offset)
}
