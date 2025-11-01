package link

import (
	"fmt"
	"net/http"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/errorz"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) Create(c *ginext.Context) {
	var body dto.CreateShortLinkRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	shortURL, err := h.service.Create(c.Request.Context(), body.URL)
	if err != nil {
		zlog.Logger.Error().
			Err(err).
			Msg("create short url")
		c.JSON(http.StatusInternalServerError, &dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.InternalServerError.Error(),
		})
		return
	}

	shortURL = fmt.Sprintf("%s/s/%s", h.publicHost, shortURL)

	resp := &dto.CreateShortLinkResponse{
		ShortURL: shortURL,
	}
	c.JSON(http.StatusCreated, &resp)
}
