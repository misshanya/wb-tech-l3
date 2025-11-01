package models

import (
	"time"

	"github.com/google/uuid"
)

type Click struct {
	ID        uuid.UUID
	LinkID    int64
	IPAddress string
	UserAgent string
	ClickedAt time.Time
}
