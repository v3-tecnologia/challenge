package services

import (
	"context"
	"fmt"
	"time"

	"bytes"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client interface {
	UploadImage(ctx context.Context, imageData []byte, deviceID string) (string, error)
	GetSignedURL(ctx context.Context, s3Key string) (string, error)
}

type S3ClientImpl struct {
	client *s3.Client
	bucket string
}

func NewS3Client() (*S3ClientImpl, error) {
	awsConfig, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsConfig)

	return &S3ClientImpl{
		client: client,
		bucket: "telemetry-photos-challenge",
	}, nil
}

func (c *S3ClientImpl) UploadImage(ctx context.Context, imageData []byte, deviceID string) (string, error) {
	timestamp := time.Now().Format("20060102150405")
	s3Key := fmt.Sprintf("photos/%s/%s.jpg", deviceID, timestamp)

	uploader := manager.NewUploader(c.client)
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(s3Key),
		Body:   bytes.NewReader(imageData),
	})

	if err != nil {
		return "", err
	}

	return s3Key, nil
}

func (c *S3ClientImpl) GetSignedURL(ctx context.Context, s3Key string) (string, error) {
	presigner := s3.NewPresignClient(c.client)

	request, err := presigner.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(s3Key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = 15 * time.Minute
	})

	if err != nil {
		return "", err
	}

	return request.URL, nil
}
