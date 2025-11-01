package link

import (
	"net/http"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/errorz"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) GetLinkStatistics(c *ginext.Context) {
	shortCode := c.Param("short")

	stats, err := h.service.GetLinkStatistics(c.Request.Context(), shortCode)
	if err != nil {
		zlog.Logger.Error().
			Err(err).
			Str("short", shortCode).
			Msg("get link statistics failed")
		c.JSON(http.StatusInternalServerError, &dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.InternalServerError.Error(),
		})
		return
	}

	userAgentsResp := make([]dto.UserAgentStats, len(stats))
	for i, stat := range stats {
		userAgentsResp[i] = dto.UserAgentStats{
			UserAgent: stat.UserAgent,
			Count:     stat.Count,
		}
	}

	resp := &dto.GetLinkStatisticsResponse{
		UserAgents: userAgentsResp,
	}
	c.JSON(http.StatusOK, resp)
}
