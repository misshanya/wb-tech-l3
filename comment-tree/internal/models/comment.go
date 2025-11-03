package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID
	Content   string
	ParentID  uuid.UUID
	Path      string
	CreatedAt time.Time
}
