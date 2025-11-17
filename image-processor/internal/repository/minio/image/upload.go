package image

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

func (r *repo) Upload(ctx context.Context, file io.Reader, size int64, filename, contentType string) error {
	_, err := r.client.PutObject(ctx, r.bucketName, filename, file, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return fmt.Errorf("failed to upload image to minio: %w", err)
	}

	return nil
}
