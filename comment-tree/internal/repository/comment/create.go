package comment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/db/sqlc/storage"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/errorz"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
)

func (r *repo) Create(ctx context.Context, c *models.Comment) (*models.Comment, error) {
	tx, err := r.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	q := r.queries.WithTx(tx)

	comment, err := q.CreateComment(ctx,
		storage.CreateCommentParams{
			Content: c.Content,
			ParentID: uuid.NullUUID{
				UUID:  c.ParentID,
				Valid: c.ParentID != uuid.Nil,
			},
		},
	)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "foreign_key_violation" {
				return nil, errorz.CommentNotFound
			}
		}
		return nil, fmt.Errorf("failed to create comment: %w", err)
	}

	var path string
	if c.ParentID != uuid.Nil {
		parentPath, err := q.GetPathByCommentID(ctx, c.ParentID)
		if err != nil {
			return nil, fmt.Errorf("failed to get parent path: %w", err)
		}
		if !parentPath.Valid {
			return nil, fmt.Errorf("parent path is not valid")
		}
		path = parentPath.String + comment.ID.String() + "/"
	} else {
		path = comment.ID.String() + "/"
	}

	err = q.UpdateCommentPath(ctx,
		storage.UpdateCommentPathParams{
			Path: sql.NullString{
				String: path,
				Valid:  true,
			},
			ID: comment.ID,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update comment path: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &models.Comment{
		ID:        comment.ID,
		Content:   c.Content,
		ParentID:  c.ParentID,
		Path:      path,
		CreatedAt: comment.CreatedAt.Time,
	}, nil
}
