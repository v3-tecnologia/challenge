package rekognition

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	rekognition "github.com/aws/aws-sdk-go-v2/service/rekognition"
	rekognitionTypes "github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/aws-sdk-go/aws"
)

type RekognitionComparer struct {
	DB *sql.DB
}

func NewRekognitionComparer(db *sql.DB) *RekognitionComparer {
	return &RekognitionComparer{DB: db}
}

func (c *RekognitionComparer) Compare(ctx context.Context, mac, filename string) (bool, error) {
	rows, err := c.DB.Query(`SELECT filename FROM photo WHERE mac = $1 AND filename != $2`, mac, filename)
	if err != nil {
		return false, custom_errors.NewDBError(err, http.StatusInternalServerError)
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

		output, err := RekognitionClient.CompareFaces(ctx, input)
		if err != nil {
			return false, custom_errors.NewRekognitionError(err, http.StatusInternalServerError)
		}

		if len(output.FaceMatches) > 0 {
			return true, nil
		}
	}
	return false, nil
}
