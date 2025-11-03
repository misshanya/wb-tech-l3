package mappers

import (
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/db/ent"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
)

func EntCommentToModel(c *ent.Comment) *models.Comment {
	return &models.Comment{
		ID:        c.ID,
		Content:   c.Content,
		ParentID:  c.ParentID,
		Path:      c.Path,
		CreatedAt: c.CreatedAt,
	}
}
