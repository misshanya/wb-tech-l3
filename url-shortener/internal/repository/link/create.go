package link

import (
	"context"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/repository/mappers"
)

func (r *repo) Create(ctx context.Context, url string) (*models.Link, error) {
	link, err := r.client.Link.
		Create().
		SetURL(url).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mappers.EntLinkToModel(link), nil
}
