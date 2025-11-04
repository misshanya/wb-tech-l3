package comment

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/errorz"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) Search(c *ginext.Context) {
	q := c.Query("q")

	comments, err := h.service.Search(c.Request.Context(), q)
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
