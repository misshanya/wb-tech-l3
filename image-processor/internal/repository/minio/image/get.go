package image

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

func (r *repo) Get(ctx context.Context, filename string) (io.ReadCloser, string, int64, error) {
	obj, err := r.client.GetObject(ctx, r.bucketName, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", 0, fmt.Errorf("failed to get image from minio: %w", err)
	}

	info, err := obj.Stat()
	if err != nil {
		return nil, "", 0, fmt.Errorf("failed to get image info from minio: %w", err)
	}

	return obj, info.ContentType, info.Size, nil
}
