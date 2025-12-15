package imageprocessor

import (
	"bytes"
	"context"
	"fmt"
	"image"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/models"
)

func (s *service) ProcessImage(ctx context.Context, id uuid.UUID) error {
	imageInfo, err := s.repo.Get(ctx, id)
	if err != nil {
		return err
	}

	// We don't need to process if status is not pending
	// However, TODO: we should handle situation when app stopped in the middle of processing
	if imageInfo.Status != models.StatusPending {
		return nil
	}

	imageData, contentType, _, err := s.imageStorage.Get(ctx, imageInfo.ID.String())
	if err != nil {
		return err
	}

	imageInfo.Status = models.StatusProcessing
	err = s.repo.UpdateStatus(ctx, imageInfo.ID, string(imageInfo.Status))
	if err != nil {
		return err
	}

	img, _, err := image.Decode(imageData)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	originalFilename := imageInfo.OriginalFilename
	format, err := imaging.FormatFromFilename(originalFilename)
	if err != nil {
		return fmt.Errorf("failed to get format from filename: %w", err)
	}

	// Process image
	var buf bytes.Buffer
	img = imaging.Resize(img, img.Bounds().Dx()/s.resizeFactor, img.Bounds().Dy()/s.resizeFactor, imaging.Lanczos)

	// Resize watermark to fit proportionally in the image
	watermark := imaging.Resize(s.watermark, img.Bounds().Dx()/2, img.Bounds().Dy()/2, imaging.Lanczos)

	img = imaging.OverlayCenter(img, watermark, 0.5)

	err = imaging.Encode(&buf, img, format)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	// Save result to the storage
	processedImageFilename := fmt.Sprintf("%s_processed", id.String())
	err = s.imageStorage.Upload(ctx, &buf, int64(buf.Len()), processedImageFilename, contentType)
	if err != nil {
		return fmt.Errorf("failed to save processed image: %w", err)
	}

	imageInfo.Status = models.StatusDone
	err = s.repo.UpdateStatus(ctx, imageInfo.ID, string(imageInfo.Status))
	if err != nil {
		return err
	}

	return nil
}
