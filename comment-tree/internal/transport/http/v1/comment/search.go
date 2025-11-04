package comment

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/errorz"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) Search(c *ginext.Context) {
	q := c.Query("q")

	var limit, offset int32
	if c.Query("limit") != "" {
		limit64, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		limit = int32(limit64)
	} else {
		limit = 20
	}

	if c.Query("page") != "" {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		offset = (int32(page) - 1) * limit
	} else {
		offset = 0
	}

	comments, err := h.service.Search(c.Request.Context(), q, limit, offset)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to search comments")
		c.JSON(http.StatusInternalServerError, &dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.InternalServerError.Error(),
		})
		return
	}

	respComments := make([]*dto.Comment, len(comments))
	for i, comment := range comments {
		respComments[i] = &dto.Comment{
			ID:        comment.ID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
		}
		if comment.ParentID != uuid.Nil {
			respComments[i].ParentID = &comment.ParentID
		}
	}
	resp := &dto.CommentsSearchResponse{
		Comments: respComments,
	}
	c.JSON(http.StatusOK, resp)
}
