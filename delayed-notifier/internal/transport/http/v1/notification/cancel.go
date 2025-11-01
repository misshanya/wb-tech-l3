package notification

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/errorz"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) Cancel(c *ginext.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = h.service.Cancel(c.Request.Context(), id)
	switch {
	case errors.Is(err, errorz.NotificationNotFound):
		zlog.Logger.Warn().Any("id", id).Msg("notification not found")
		c.JSON(http.StatusNotFound, &dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	case err != nil:
		zlog.Logger.Error().Any("id", id).Msg("failed to cancel")
		c.JSON(http.StatusInternalServerError, &dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.InternalServerError.Error(),
		})
		return
	}
}
