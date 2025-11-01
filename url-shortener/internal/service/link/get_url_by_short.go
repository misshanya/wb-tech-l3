package link

import (
	"context"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
	"github.com/misshanya/wb-tech-l3/url-shortener/pkg/base62"
	"github.com/wb-go/wbf/zlog"
)

func (s *service) GetURLByShort(ctx context.Context, short string, clickInfo *models.Click) (string, error) {
	id := base62.Decode(short)
	link, err := s.repo.Get(ctx, id)
	if err != nil {
		return "", err
	}

	clickInfo.LinkID = link.ID

	go func() {
		_, err = s.repo.AddClick(context.Background(), clickInfo)
		if err != nil {
			zlog.Logger.Error().
				Err(err).
				Msg("failed to add click")
		}
	}()

	return link.URL, nil
}
