package link

import (
	"context"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/db/ent"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/db/ent/link"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/errorz"
)

func (r *repo) GetIDByURL(ctx context.Context, url string) (int64, error) {
	l, err := r.client.Link.
		Query().
		Where(link.URL(url)).
		Select(link.FieldID).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return 0, errorz.LinkNotFound
		}
		return 0, err
	}
	return l.ID, nil
}
