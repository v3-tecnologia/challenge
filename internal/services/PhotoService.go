package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"time"

	"github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/db"
	"github.com/KaiRibeiro/challenge/internal/models"
	s3Config "github.com/KaiRibeiro/challenge/internal/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func AddPhoto(photo models.PhotoModel) error {
	ctx := context.Background()

	timestamp := time.UnixMilli(photo.Timestamp)

	imageBytes, err := base64.StdEncoding.DecodeString(photo.ImageBase64)
	if err != nil {
		return errors.New("Failed to decode base64 image: " + err.Error())
	}

	filename := timestamp.Format("20060102150405") + "-photo.jpg"

	_, err = s3Config.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &config.BucketName,
		Key:         awsString(filename),
		Body:        bytes.NewReader(imageBytes),
		ContentType: awsString("image/jpeg"),
	})
	if err != nil {
		return errors.New("Failed to upload image to s3: " + err.Error())
	}

	file_url := "https://" + config.BucketName + ".s3." + config.AwsRegion + ".amazonaws.com/" + filename

	query := `INSERT INTO photo (filename, file_url, mac, timestamp)
    VALUES ($1, $2, $3, $4)`

	_, err = db.DB.Exec(query, filename, file_url, photo.MAC, timestamp)

	return err
}

func awsString(s string) *string { return &s }
