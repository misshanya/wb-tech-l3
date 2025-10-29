package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID          uuid.UUID
	ScheduledAt time.Time
	Title       string
	Content     string
	Channel     string
	Receiver    string
}
