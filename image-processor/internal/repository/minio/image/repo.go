package image

import (
	"github.com/minio/minio-go/v7"
)

type repo struct {
	client     *minio.Client
	bucketName string
}

func New(client *minio.Client, bucketName string) *repo {
	return &repo{
		client:     client,
		bucketName: bucketName,
	}
}
