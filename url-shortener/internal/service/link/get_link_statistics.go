package link

import (
	"context"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
	"github.com/misshanya/wb-tech-l3/url-shortener/pkg/base62"
)

func (s *service) GetLinkStatistics(ctx context.Context, short string) (models.LinkStatistics, error) {
	id := base62.Decode(short)
	return s.repo.GetLinkStatistics(ctx, id)
}
