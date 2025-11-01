package link

import (
	"context"

	"github.com/wb-go/wbf/ginext"
)

type service interface {
	Create(ctx context.Context, url string) (string, error)
	GetURLByShort(ctx context.Context, short string) (string, error)
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
}
