package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationStatus string

var (
	StatusScheduled  NotificationStatus = "scheduled"
	StatusProcessing NotificationStatus = "processing"
	StatusDelivered  NotificationStatus = "delivered"
	StatusCancelled  NotificationStatus = "cancelled"
	StatusFailed     NotificationStatus = "failed"
)

type Notification struct {
	ID          uuid.UUID
	ScheduledAt time.Time
	Title       string
	Content     string
	Channel     string
	Receiver    string
	Status      NotificationStatus
}
