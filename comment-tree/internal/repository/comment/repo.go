package comment

import "github.com/misshanya/wb-tech-l3/comment-tree/internal/db/ent"

type repo struct {
	client *ent.Client
}

func New(client *ent.Client) *repo {
	return &repo{client: client}
}
