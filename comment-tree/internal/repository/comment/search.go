package comment

import (
	"context"
	"fmt"

	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
)

func (r *repo) Search(ctx context.Context, query string) ([]*models.Comment, error) {
	comments, err := r.queries.SearchComments(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to search comments: %w", err)
	}

	res := make([]*models.Comment, len(comments))
	for i, c := range comments {
		res[i] = &models.Comment{
			ID:        c.ID,
			Content:   c.Content,
			ParentID:  c.ParentID.UUID,
			Path:      c.Path.String,
			CreatedAt: c.CreatedAt.Time,
		}
	}

	return res, nil
}
