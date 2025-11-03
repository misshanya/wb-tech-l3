package comment

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/db/ent"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/db/ent/comment"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/errorz"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/repository/mappers"
)

func (r *repo) GetDerivatives(ctx context.Context, id uuid.UUID) ([]*models.Comment, error) {
	parent, err := r.client.Comment.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorz.CommentNotFound
		}
		return nil, fmt.Errorf("failed to get parent: %w", err)
	}

	derivatives, err := r.client.Comment.Query().
		Where(comment.PathHasPrefix(parent.Path)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get derivatives: %w", err)
	}

	res := make([]*models.Comment, len(derivatives))
	for i, d := range derivatives {
		res[i] = mappers.EntCommentToModel(d)
	}

	return res, nil
}
