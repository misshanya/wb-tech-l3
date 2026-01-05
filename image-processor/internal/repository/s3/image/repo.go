package image

import (
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type repo struct {
	client     *s3.Client
	bucketName string
	uploader   *manager.Uploader
}

func New(client *s3.Client, bucketName string) *repo {
	return &repo{
		client:     client,
		bucketName: bucketName,
		uploader:   manager.NewUploader(client),
	}
}
