package comment

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/db/ent/comment"
)

func (r *repo) Delete(ctx context.Context, id uuid.UUID) error {
	parent, err := r.client.Comment.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get parent: %w", err)
	}

	_, err = r.client.Comment.Delete().Where(comment.PathHasPrefix(parent.Path)).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}

	return nil
}
