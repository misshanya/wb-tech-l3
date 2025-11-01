package link

import (
	"context"

	"github.com/misshanya/wb-tech-l3/url-shortener/pkg/base62"
)

func (s *service) GetURLByShort(ctx context.Context, short string) (string, error) {
	id := base62.Decode(short)
	link, err := s.repo.Get(ctx, id)
	if err != nil {
		return "", err
	}
	return link.URL, nil
}
