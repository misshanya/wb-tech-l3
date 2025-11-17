package image

import (
	"context"
	"io"

	"github.com/misshanya/wb-tech-l3/image-processor/internal/models"
	"github.com/wb-go/wbf/ginext"
)

type service interface {
	Upload(ctx context.Context, content io.Reader, size int64, filename, contentType string) (*models.Image, error)
}

type handler struct {
	service service
}

func New(s service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) Setup(group *ginext.RouterGroup) {
	group.POST("/upload", h.Upload)
}
