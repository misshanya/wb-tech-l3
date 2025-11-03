package comment

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/db/ent"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/errorz"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/repository/mappers"
)

func (r *repo) Create(ctx context.Context, c *models.Comment) (*models.Comment, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	commentCreate := tx.Comment.
		Create().
		SetContent(c.Content)

	if c.ParentID != uuid.Nil {
		commentCreate = commentCreate.SetParentID(c.ParentID)
	}

	comment, err := commentCreate.Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, errorz.CommentNotFound
		}
		return nil, fmt.Errorf("failed to save comment: %w", err)
	}

	var path string
	if c.ParentID != uuid.Nil {
		parent, err := tx.Comment.Get(ctx, c.ParentID)
		if err != nil {
			return nil, fmt.Errorf("failed to get parent comment: %w", err)
		}
		path = parent.Path + comment.ID.String() + "/"
	} else {
		path = comment.ID.String() + "/"
	}

	comment, err = tx.Comment.
		UpdateOneID(comment.ID).
		SetPath(path).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update comment to set path: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return mappers.EntCommentToModel(comment), nil
}
