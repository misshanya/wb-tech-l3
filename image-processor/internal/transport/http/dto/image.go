package dto

import "github.com/google/uuid"

type ImageInfo struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

type ImageUploadResponse ImageInfo
