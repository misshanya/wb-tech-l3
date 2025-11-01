package link

import (
	"errors"
	"net/http"
	"time"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/errorz"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) Redirect(c *ginext.Context) {
	shortCode := c.Param("short")

	clickInfo := &models.Click{
		IPAddress: c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		ClickedAt: time.Now(),
	}
	url, err := h.service.GetURLByShort(c.Request.Context(), shortCode, clickInfo)
	switch {
	case errors.Is(err, errorz.LinkNotFound):
		zlog.Logger.Warn().
			Str("short", shortCode).
			Msg("url not found")
		c.JSON(http.StatusNotFound, &dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	case err != nil:
		zlog.Logger.Error().
			Err(err).
			Msg("failed to get url")
		c.JSON(http.StatusInternalServerError, &dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.InternalServerError.Error(),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
