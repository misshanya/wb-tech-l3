package comment

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/errorz"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) Get(c *ginext.Context) {
	idStr := c.Query("parent")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

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

	if limit < 0 {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprint("limit must be greater or equal than 0"),
		})
		return
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

	if offset < 0 {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprint("page must be greater or equal than 1"),
		})
		return
	}

	comments, err := h.service.Get(c.Request.Context(), id, limit, offset)
	if err != nil {
		if errors.Is(err, errorz.CommentNotFound) {
			zlog.Logger.Warn().Any("id", id).Msg("comment not found")
			c.JSON(http.StatusNotFound, &dto.HTTPStatus{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
			return
		}
		zlog.Logger.Error().Err(err).Msg("failed to get comment")
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
	resp := &dto.CommentsGetResponse{
		Comments: respComments,
	}
	c.JSON(http.StatusOK, resp)
}
