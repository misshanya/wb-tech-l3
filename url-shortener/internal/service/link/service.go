package link

import (
	"context"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
)

type repo interface {
	Create(ctx context.Context, url string) (*models.Link, error)
	Get(ctx context.Context, id int64) (*models.Link, error)
	GetIDByURL(ctx context.Context, url string) (int64, error)
}

type service struct {
	repo repo
}

func New(repo repo) *service {
	return &service{repo: repo}
}
