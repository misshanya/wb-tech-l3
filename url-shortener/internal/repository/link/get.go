package link

import (
	"context"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/db/ent/link"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/repository/mappers"
)

func (r *repo) Get(ctx context.Context, id int64) (*models.Link, error) {
	l, err := r.client.Link.
		Query().
		Where(link.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return mappers.EntLinkToModel(l), nil
}
