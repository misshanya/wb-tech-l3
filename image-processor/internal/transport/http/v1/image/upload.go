package image

import (
	"bytes"
	"io"
	"net/http"

	"github.com/misshanya/wb-tech-l3/image-processor/internal/errorz"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/transport/http/dto"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *handler) Upload(c *ginext.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: "file is required",
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to open file")
		c.JSON(http.StatusInternalServerError, &dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.ErrInternalServerError.Error(),
		})
		return
	}
	defer src.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, src); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to read file")
		c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.ErrInternalServerError.Error(),
		})
		return
	}

	imageInfo, err := h.service.Upload(c.Request.Context(), bytes.NewReader(buf.Bytes()), file.Size, file.Filename, file.Header.Get("Content-Type"))
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to upload image")
		c.JSON(http.StatusInternalServerError, &dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.ErrInternalServerError.Error(),
		})
		return
	}

	resp := &dto.ImageUploadResponse{
		ID:     imageInfo.ID,
		Status: string(imageInfo.Status),
	}
	c.JSON(http.StatusAccepted, resp)
}
