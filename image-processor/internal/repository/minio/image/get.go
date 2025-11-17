package image

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

func (r *repo) Get(ctx context.Context, filename string) (io.Reader, string, int64, error) {
	obj, err := r.client.GetObject(ctx, r.bucketName, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", 0, fmt.Errorf("failed to get image from minio: %w", err)
	}
	defer obj.Close()

	info, err := obj.Stat()
	if err != nil {
		return nil, "", 0, fmt.Errorf("failed to get image info from minio: %w", err)
	}

	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, "", 0, fmt.Errorf("failed to read image data from minio: %w", err)
	}

	return bytes.NewReader(data), info.ContentType, int64(len(data)), nil
}
