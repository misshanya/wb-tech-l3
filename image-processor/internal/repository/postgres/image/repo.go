package image

import "github.com/misshanya/wb-tech-l3/image-processor/internal/db/sqlc/storage"

type repo struct {
	queries *storage.Queries
}

func New(queries *storage.Queries) *repo {
	return &repo{queries: queries}
}
