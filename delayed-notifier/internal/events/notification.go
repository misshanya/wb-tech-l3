package events

import "github.com/google/uuid"

type Notification struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Channel  string    `json:"channel"`
	Receiver string    `json:"receiver"`
}
