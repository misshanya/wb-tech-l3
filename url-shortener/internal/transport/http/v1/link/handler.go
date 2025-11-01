package link

import (
	"context"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
	"github.com/wb-go/wbf/ginext"
)

type service interface {
	Create(ctx context.Context, url string) (string, error)
	GetURLByShort(ctx context.Context, short string, clickInfo *models.Click) (string, error)
	GetLinkStatistics(ctx context.Context, short string) (models.LinkStatistics, error)
}

type handler struct {
	service    service
	publicHost string
}

func New(service service, publicHost string) *handler {
	return &handler{service: service, publicHost: publicHost}
}

func (h *handler) Setup(group *ginext.RouterGroup) {
	group.POST("/shorten", h.Create)
	group.GET("/s/:short", h.Redirect)
	group.GET("/analytics/:short", h.GetLinkStatistics)
}
