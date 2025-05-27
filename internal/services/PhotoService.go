package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/db"
	"github.com/KaiRibeiro/challenge/internal/models"
	rekognitionConfig "github.com/KaiRibeiro/challenge/internal/rekognition"
	s3Config "github.com/KaiRibeiro/challenge/internal/s3"
	"github.com/aws/aws-sdk-go-v2/aws"
	rekognition "github.com/aws/aws-sdk-go-v2/service/rekognition"
	rekognitionTypes "github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func AddPhoto(photo models.PhotoModel) (bool, error) {
	ctx := context.Background()

	timestamp := time.UnixMilli(photo.Timestamp)

	imageBytes, err := base64.StdEncoding.DecodeString(photo.ImageBase64)
	if err != nil {
		return false, errors.New("Failed to decode base64 image: " + err.Error())
	}

	filename := fmt.Sprintf("%s-%d-photo.jpg", timestamp.Format("20060102150405"), time.Now().UnixNano())

	_, err = s3Config.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &config.BucketName,
		Key:         awsString(filename),
		Body:        bytes.NewReader(imageBytes),
		ContentType: awsString("image/jpeg"),
	})
	if err != nil {
		return false, errors.New("Failed to upload image to s3: " + err.Error())
	}

	file_url := "https://" + config.BucketName + ".s3." + config.AwsRegion + ".amazonaws.com/" + filename

	query := `INSERT INTO photo (filename, file_url, mac, timestamp)
    VALUES ($1, $2, $3, $4)`

	_, err = db.DB.Exec(query, filename, file_url, photo.MAC, timestamp)

	recognized, err := CompareWithPreviousFaces(ctx, photo.MAC, filename)

	return recognized, err
}

func CompareWithPreviousFaces(ctx context.Context, mac string, filename string) (bool, error) {
	rows, err := db.DB.Query(`SELECT filename FROM photo WHERE mac = $1 AND filename != $2`, mac, filename)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var refFilename string
		if err := rows.Scan(&refFilename); err != nil {
			continue
		}

		input := &rekognition.CompareFacesInput{
			SourceImage: &rekognitionTypes.Image{
				S3Object: &rekognitionTypes.S3Object{
					Bucket: &config.BucketName,
					Name:   &refFilename,
				},
			},
			TargetImage: &rekognitionTypes.Image{
				S3Object: &rekognitionTypes.S3Object{
					Bucket: &config.BucketName,
					Name:   &filename,
				},
			},
			SimilarityThreshold: aws.Float32(90.0),
		}

		output, err := rekognitionConfig.RekognitionClient.CompareFaces(ctx, input)
		if err != nil {
			fmt.Printf("Rekognition error: %v\n", err)
			continue
		}

		if len(output.FaceMatches) > 0 {
			fmt.Printf("Face match found: %+v\n", output.FaceMatches)
			return true, nil
		} else {
			fmt.Println("No face match for this pair.")
		}
	}

	return false, nil
}

func awsString(s string) *string { return &s }
