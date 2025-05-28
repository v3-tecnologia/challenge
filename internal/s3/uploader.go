package s3

import (
	"bytes"
	"context"
	"net/http"

	"github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Uploader struct{}

func (u *S3Uploader) PutPhoto(ctx context.Context, filename string, image []byte) (string, error) {
	_, err := S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &config.BucketName,
		Key:         aws.String(filename),
		Body:        bytes.NewReader(image),
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return "", custom_errors.NewS3Error(err, http.StatusInternalServerError)
	}
	url := "https://" + config.BucketName + ".s3." + config.AwsRegion + ".amazonaws.com/" + filename
	return url, nil
}
