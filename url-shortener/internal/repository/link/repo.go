package link

import "github.com/misshanya/wb-tech-l3/url-shortener/internal/db/ent"

type repo struct {
	client *ent.Client
}

func New(client *ent.Client) *repo {
	return &repo{client: client}
}
