package notification

import "github.com/misshanya/wb-tech-l3/delayed-notifier/internal/db/ent"

type repo struct {
	client *ent.Client
}

func New(client *ent.Client) *repo {
	return &repo{client: client}
}
