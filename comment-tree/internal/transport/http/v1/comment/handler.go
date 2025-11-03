package comment

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
	"github.com/wb-go/wbf/ginext"
)

type service interface {
	Create(ctx context.Context, c *models.Comment) (*models.Comment, error)
	Get(ctx context.Context, id uuid.UUID) ([]*models.Comment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type handler struct {
	service   service
	validator *validator.Validate
}

func New(s service) *handler {
	return &handler{
		s,
		validator.New(),
	}
}

func (h *handler) Setup(group *ginext.RouterGroup) {
	group.POST("/", h.Create)
	group.GET("/", h.Get)
	group.DELETE("/:id", h.Delete)
}
