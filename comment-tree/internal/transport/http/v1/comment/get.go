package comment

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/errorz"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) Get(c *ginext.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	comments, err := h.service.Get(c.Request.Context(), id)
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
