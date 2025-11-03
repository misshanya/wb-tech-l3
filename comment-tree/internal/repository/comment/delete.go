package comment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/errorz"
)

func (r *repo) Delete(ctx context.Context, id uuid.UUID) error {
	parentPath, err := r.queries.GetPathByCommentID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorz.CommentNotFound
		}
		return fmt.Errorf("failed to get parent path: %w", err)
	}

	if !parentPath.Valid {
		return errorz.CommentNotFound
	}

	err = r.queries.DeleteComment(ctx, parentPath.String)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}

	return nil
}
