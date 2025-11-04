package dto

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID  `json:"id"`
	Content   string     `json:"content"`
	ParentID  *uuid.UUID `json:"parent_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

type CommentCreateRequest struct {
	Content  string    `json:"content" validate:"required,min=1,max=1500"`
	ParentID uuid.UUID `json:"parent_id,omitempty" validate:"omitempty"`
}

type CommentCreateResponse Comment

type CommentsGetResponse struct {
	Comments []*Comment `json:"comments"`
}

type CommentsSearchResponse struct {
	Comments []*Comment `json:"comments"`
}
