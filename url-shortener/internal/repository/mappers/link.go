package mappers

import (
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/db/ent"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
)

func EntLinkToModel(l *ent.Link) *models.Link {
	return &models.Link{
		ID:  l.ID,
		URL: l.URL,
	}
}
