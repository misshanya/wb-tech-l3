package notification

import (
	"github.com/go-playground/validator/v10"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
	"github.com/wb-go/wbf/ginext"
)

type service interface {
	Create(n *models.Notification) (*models.Notification, error)
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
}
