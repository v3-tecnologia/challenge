package services

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/KaiRibeiro/challenge/internal/interfaces"
	"github.com/KaiRibeiro/challenge/internal/logs"
	"github.com/KaiRibeiro/challenge/internal/models"
)

type PhotoService interface {
	AddPhoto(photo models.PhotoModel) (bool, error)
}

type PhotoDBService struct {
	DB           *sql.DB
	Uploader     interfaces.Uploader
	FaceComparer interfaces.FaceComparer
}

func NewPhotoDBService(dbConn *sql.DB, uploader interfaces.Uploader, comparer interfaces.FaceComparer) *PhotoDBService {
	return &PhotoDBService{
		DB:           dbConn,
		Uploader:     uploader,
		FaceComparer: comparer,
	}
}

func (s *PhotoDBService) AddPhoto(photo models.PhotoModel) (bool, error) {
	logs.Logger.Info("adding photo data to database",
		"mac", photo.MAC,
		"timestamp", photo.Timestamp,
		"image_base64_length", len(photo.ImageBase64),
	)
	ctx := context.Background()

	timestamp := time.UnixMilli(photo.Timestamp)

	logs.Logger.Info("decoding base64 image",
		"image_base64_length", len(photo.ImageBase64),
	)
	imageBytes, err := base64.StdEncoding.DecodeString(photo.ImageBase64)
	if err != nil {
		wrappedErr := fmt.Errorf("failed to decode base64: %w", custom_errors.NewPhotoError(err, http.StatusInternalServerError))
		logs.Logger.Error("failed to decode base64 image",
			"error", wrappedErr,
		)
		return false, wrappedErr
	}

	filename := fmt.Sprintf("%s-%d-photo.jpg", timestamp.Format("20060102150405"), time.Now().UnixNano())

	fileURL, err := s.Uploader.PutPhoto(ctx, filename, imageBytes)
	if err != nil {
		wrappedErr := fmt.Errorf("failed to upload image to s3: %w", err)
		return false, wrappedErr
	}

	query := `INSERT INTO photo (filename, file_url, mac, timestamp)
    VALUES ($1, $2, $3, $4)`

	_, err = s.DB.Exec(query, filename, fileURL, photo.MAC, timestamp)
	if err != nil {
		wrappedErr := fmt.Errorf("failed to insert photo data into database: %w", custom_errors.NewDBError(err, http.StatusInternalServerError))
		logs.Logger.Error("failed to add photo data",
			"error", wrappedErr,
		)
		return false, wrappedErr
	}

	logs.Logger.Info("Comparing image",
		"url", fileURL,
	)
	recognized, err := s.FaceComparer.Compare(ctx, photo.MAC, filename)
	if err != nil {
		wrappedErr := fmt.Errorf("failed to compare image: %w", err)
		return false, wrappedErr
	}

	return recognized, err
}
