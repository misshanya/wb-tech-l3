package comment

import (
	"context"

	"github.com/google/uuid"
)

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
