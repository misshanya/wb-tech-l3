package image

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (r *repo) Get(ctx context.Context, filename string) (io.ReadCloser, string, int64, error) {
	obj, err := r.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, "", 0, fmt.Errorf("failed to get image from s3: %w", err)
	}

	return obj.Body, *obj.ContentType, *obj.ContentLength, nil
}
