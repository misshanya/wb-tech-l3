package image

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/errorz"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) GetProcessed(c *ginext.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	image, contentType, status, size, err := h.service.GetProcessed(c.Request.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, errorz.ErrImageIsNotDone):
			c.JSON(http.StatusAccepted, &dto.ImageInfo{
				ID:     id,
				Status: string(status),
			})
		case errors.Is(err, errorz.ErrImageNotFound):
			c.JSON(http.StatusNotFound, &dto.HTTPStatus{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		default:
			zlog.Logger.Error().Err(err).Msg("failed to get image")
			c.JSON(http.StatusInternalServerError, &dto.HTTPStatus{
				Code:    http.StatusInternalServerError,
				Message: errorz.ErrInternalServerError.Error(),
			})
		}
		return
	}

	c.DataFromReader(
		http.StatusOK,
		size,
		contentType,
		image,
		nil,
	)
}
