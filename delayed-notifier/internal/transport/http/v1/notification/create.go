package notification

import (
	"net/http"
	"time"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/errorz"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) Create(c *ginext.Context) {
	var body dto.NotificationCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := h.validator.StructCtx(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if body.ScheduledAt.IsZero() {
		body.ScheduledAt = time.Now()
	}

	n := &models.Notification{
		ScheduledAt: body.ScheduledAt,
		Title:       body.Title,
		Content:     body.Content,
		Channel:     body.Channel,
		Receiver:    body.Receiver,
	}
	notification, err := h.service.Create(c.Request.Context(), n)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("create notification")
		c.JSON(http.StatusInternalServerError, &dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.InternalServerError.Error(),
		})
		return
	}

	zlog.Logger.Info().Any("notification", notification).Msg("create notification")

	resp := &dto.NotificationCreateResponse{
		ID:          notification.ID,
		ScheduledAt: notification.ScheduledAt,
		Title:       notification.Title,
		Content:     notification.Content,
		Channel:     notification.Channel,
		Receiver:    notification.Receiver,
		Status:      string(notification.Status),
	}
	c.JSON(http.StatusCreated, resp)
}
