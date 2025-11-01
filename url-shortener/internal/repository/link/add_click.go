package link

import (
	"context"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/repository/mappers"
)

func (r *repo) AddClick(ctx context.Context, click *models.Click) (*models.Click, error) {
	c, err := r.client.Click.
		Create().
		SetLinkID(click.LinkID).
		SetIP(click.IPAddress).
		SetUserAgent(click.UserAgent).
		SetClickedAt(click.ClickedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	model := mappers.EntClickToModel(c)
	model.LinkID = click.LinkID
	return model, nil
}
