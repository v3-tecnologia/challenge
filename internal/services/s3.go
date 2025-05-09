package services

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ricardoraposo/challenge/internal/interfaces"
)

var (
	s3Region           string = os.Getenv("AWS_REGION")
	awsAccessKeyID     string = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey string = os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsConfig                 = aws.Config{
		Region:      s3Region,
		Credentials: credentials.NewStaticCredentialsProvider(awsAccessKeyID, awsSecretAccessKey, ""),
	}
)

type S3Uploader struct {
	client *s3.Client
}

func NewS3Uploader() interfaces.BucketUploader {
	client := s3.NewFromConfig(awsConfig)

	return &S3Uploader{
		client: client,
	}
}

func (up *S3Uploader) UploadAsync(file io.Reader, key string, ch chan<- string) {
	uploader := manager.NewUploader(up.client)
	uploadedFile, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("AWS_REGION")),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String("image/png"),
		ACL:         "public-read",
	})

	if err != nil {
		ch <- "https://i.imgur.com/fLjGgnc.png"
	}

	ch <- uploadedFile.Location
}
