package comment

import (
	"database/sql"

	"github.com/misshanya/wb-tech-l3/comment-tree/internal/db/sqlc/storage"
)

type repo struct {
	dbConn  *sql.DB
	queries *storage.Queries
}

func New(dbConn *sql.DB, queries *storage.Queries) *repo {
	return &repo{dbConn: dbConn, queries: queries}
}
