package dto

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID          uuid.UUID `json:"id"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Channel     string    `json:"channel"`
	Receiver    string    `json:"receiver"`
}

type NotificationCreateRequest struct {
	ScheduledAt time.Time `json:"scheduled_at,omitempty" validate:"omitempty"`
	Title       string    `json:"title" validate:"required"`
	Content     string    `json:"content" validate:"required"`
	Channel     string    `json:"channel" validate:"required"`
	Receiver    string    `json:"receiver" validate:"required"`
}

type NotificationCreateResponse Notification
