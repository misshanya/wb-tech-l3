package comment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/errorz"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
)

func (r *repo) GetDerivatives(ctx context.Context, id uuid.UUID) ([]*models.Comment, error) {
	parentPath, err := r.queries.GetPathByCommentID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*models.Comment{}, errorz.CommentNotFound
		}
		return []*models.Comment{}, fmt.Errorf("failed to get parent path: %w", err)
	}

	if !parentPath.Valid {
		return []*models.Comment{}, errorz.CommentNotFound
	}

	derivatives, err := r.queries.GetDerivatives(ctx, parentPath.String)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*models.Comment{}, errorz.CommentNotFound
		}
		return []*models.Comment{}, fmt.Errorf("failed to get derivatives: %w", err)
	}

	res := make([]*models.Comment, len(derivatives))
	for i, d := range derivatives {
		res[i] = &models.Comment{
			ID:        d.ID,
			Content:   d.Content,
			ParentID:  d.ParentID.UUID,
			Path:      d.Path.String,
			CreatedAt: d.CreatedAt.Time,
		}
	}

	return res, nil
}
