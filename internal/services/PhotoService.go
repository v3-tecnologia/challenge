package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/KaiRibeiro/challenge/internal/db"
	"github.com/KaiRibeiro/challenge/internal/models"
	rekognitionConfig "github.com/KaiRibeiro/challenge/internal/rekognition"
	s3Config "github.com/KaiRibeiro/challenge/internal/s3"
	"github.com/aws/aws-sdk-go-v2/aws"
	rekognition "github.com/aws/aws-sdk-go-v2/service/rekognition"
	rekognitionTypes "github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type PhotoService interface {
	AddPhoto(photo models.PhotoModel) (bool, error)
}

type PhotoDBService struct {
	DB *sql.DB
}

func NewPhotoDBService(dbConn *sql.DB) *PhotoDBService {
	return &PhotoDBService{DB: dbConn}
}

func (s *PhotoDBService) AddPhoto(photo models.PhotoModel) (bool, error) {
	ctx := context.Background()

	timestamp := time.UnixMilli(photo.Timestamp)

	imageBytes, err := base64.StdEncoding.DecodeString(photo.ImageBase64)
	if err != nil {
		return false, fmt.Errorf("failed to decode base64: %w", custom_errors.NewPhotoError(err, http.StatusInternalServerError))
	}

	filename := fmt.Sprintf("%s-%d-photo.jpg", timestamp.Format("20060102150405"), time.Now().UnixNano())

	_, err = s3Config.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &config.BucketName,
		Key:         awsString(filename),
		Body:        bytes.NewReader(imageBytes),
		ContentType: awsString("image/jpeg"),
	})
	if err != nil {
		return false, fmt.Errorf("failed to upload image to s3: %w", custom_errors.NewS3Error(err, http.StatusInternalServerError))
	}

	file_url := "https://" + config.BucketName + ".s3." + config.AwsRegion + ".amazonaws.com/" + filename

	query := `INSERT INTO photo (filename, file_url, mac, timestamp)
    VALUES ($1, $2, $3, $4)`

	_, err = s.DB.Exec(query, filename, file_url, photo.MAC, timestamp)
	if err != nil {
		return false, fmt.Errorf("failed to insert photo data into database: %w", custom_errors.NewDBError(err, http.StatusInternalServerError))
	}
	recognized, err := CompareWithPreviousFaces(ctx, photo.MAC, filename)

	return recognized, err
}

func CompareWithPreviousFaces(ctx context.Context, mac string, filename string) (bool, error) {
	rows, err := db.DB.Query(`SELECT filename FROM photo WHERE mac = $1 AND filename != $2`, mac, filename)
	if err != nil {
		return false, fmt.Errorf("failed to find previous photos: %w", custom_errors.NewDBError(err, http.StatusInternalServerError))
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
			return false, fmt.Errorf("rekognition error: %w", custom_errors.NewRekognitionError(err, http.StatusInternalServerError))
		}

		if len(output.FaceMatches) > 0 {
			return true, nil
		}
	}

	return false, err
}

func awsString(s string) *string { return &s }
