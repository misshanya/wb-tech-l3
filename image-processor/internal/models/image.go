package models

import "github.com/google/uuid"

type Status string

var (
	StatusUploading  Status = "uploading"
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusDone       Status = "done"
)

type Image struct {
	ID               uuid.UUID
	OriginalFilename string
	Status           Status
}
