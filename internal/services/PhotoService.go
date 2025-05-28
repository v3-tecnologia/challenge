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
	ctx := context.Background()

	timestamp := time.UnixMilli(photo.Timestamp)

	imageBytes, err := base64.StdEncoding.DecodeString(photo.ImageBase64)
	if err != nil {
		return false, fmt.Errorf("failed to decode base64: %w", custom_errors.NewPhotoError(err, http.StatusInternalServerError))
	}

	filename := fmt.Sprintf("%s-%d-photo.jpg", timestamp.Format("20060102150405"), time.Now().UnixNano())

	fileURL, err := s.Uploader.PutPhoto(ctx, filename, imageBytes)
	if err != nil {
		return false, fmt.Errorf("failed to upload image: %w", err)
	}

	query := `INSERT INTO photo (filename, file_url, mac, timestamp)
    VALUES ($1, $2, $3, $4)`

	_, err = s.DB.Exec(query, filename, fileURL, photo.MAC, timestamp)
	if err != nil {
		return false, fmt.Errorf("failed to insert photo data into database: %w", custom_errors.NewDBError(err, http.StatusInternalServerError))
	}
	recognized, err := s.FaceComparer.Compare(ctx, photo.MAC, filename)

	return recognized, err
}
