package mappers

import (
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/db/ent"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
)

func EntClickToModel(click *ent.Click) *models.Click {
	return &models.Click{
		ID:        click.ID,
		IPAddress: click.IP,
		UserAgent: click.UserAgent,
		ClickedAt: click.ClickedAt,
	}
}

func EntClickWithLinkToModel(click *ent.Click) *models.Click {
	return &models.Click{
		ID:        click.ID,
		LinkID:    click.Edges.Link.ID,
		IPAddress: click.IP,
		UserAgent: click.UserAgent,
		ClickedAt: click.ClickedAt,
	}
}
