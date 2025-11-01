package notification

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
	"github.com/wb-go/wbf/ginext"
)

type service interface {
	Create(ctx context.Context, n *models.Notification) (*models.Notification, error)
	Get(ctx context.Context, id uuid.UUID) (*models.Notification, error)
	Cancel(ctx context.Context, id uuid.UUID) error
}

type handler struct {
	service   service
	validator *validator.Validate
}

func New(service service) *handler {
	return &handler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *handler) Setup(group *ginext.RouterGroup) {
	group.POST("/", h.Create)
	group.GET("/:id", h.Get)
	group.DELETE("/:id", h.Cancel)
}
