package services

import (
	"github.com/KaiRibeiro/challenge/internal/db"
	"github.com/KaiRibeiro/challenge/internal/models"
	"time"
)

func AddPhoto(photo models.PhotoModel) error {
	timestamp := time.UnixMilli(photo.Timestamp)

	query := `INSERT INTO photo (image_base64, mac, timestamp)
    VALUES ($1, $2, $3)`

	_, err := db.DB.Exec(query, photo.ImageBase64, photo.MAC, timestamp)

	return err
}
