package image

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/misshanya/wb-tech-l3/image-processor/internal/errorz"
)

func (r *repo) Get(ctx context.Context, filename string) (io.ReadCloser, string, int64, error) {
	obj, err := r.client.GetObject(ctx, r.bucketName, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", 0, fmt.Errorf("failed to get image from minio: %w", err)
	}

	info, err := obj.Stat()
	if err != nil {
		if err := obj.Close(); err != nil {
			return nil, "", 0, fmt.Errorf("failed to close minio object: %w", err)
		}
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return nil, "", 0, errorz.ErrImageNotFound
		}
		return nil, "", 0, fmt.Errorf("failed to get image info from minio: %w", err)
	}

	return obj, info.ContentType, info.Size, nil
}
