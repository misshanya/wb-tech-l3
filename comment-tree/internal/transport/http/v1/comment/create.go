package comment

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/errorz"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/models"
	"github.com/misshanya/wb-tech-l3/comment-tree/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) Create(c *ginext.Context) {
	var body dto.CommentCreateRequest
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	commentInput := &models.Comment{
		Content:  body.Content,
		ParentID: body.ParentID,
	}
	comment, err := h.service.Create(c.Request.Context(), commentInput)
	if err != nil {
		if errors.Is(err, errorz.CommentNotFound) {
			c.JSON(http.StatusNotFound, &dto.HTTPStatus{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
			return
		}
		zlog.Logger.Error().Err(err).Msg("failed to create comment")
		c.JSON(http.StatusInternalServerError, &dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.InternalServerError.Error(),
		})
		return
	}

	resp := &dto.CommentCreateResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}
	if comment.ParentID != uuid.Nil {
		resp.ParentID = &comment.ParentID
	}
	c.JSON(http.StatusCreated, resp)
}
