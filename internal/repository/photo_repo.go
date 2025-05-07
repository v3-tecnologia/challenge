package repository

import (
	"database/sql"

	"github.com/bielgennaro/v3-challenge-cloud/internal/db"
	"github.com/bielgennaro/v3-challenge-cloud/internal/model"
	"github.com/google/uuid"
)

func SavePhoto(photo *model.Photo) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}

	// Insertion Query
	query := `
        INSERT INTO photo_data (id, mac_address, file_path, recognized, timestamp, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `

	_, err = conn.Exec(
		query,
		photo.ID,
		photo.MacAddress,
		photo.FileURL,
		photo.IsMatch,
		photo.Timestamp,
		photo.CreatedAt,
	)

	return err
}

func GetPhotoByID(id uuid.UUID) (*model.Photo, error) {
	conn, err := db.GetConnection()
	if err != nil {
		return nil, err
	}

	query := `
        SELECT id, mac_address, file_path, recognized, timestamp, created_at
        FROM photo_data
        WHERE id = $1
    `

	var photo model.Photo
	err = conn.QueryRow(query, id).Scan(
		&photo.ID,
		&photo.MacAddress,
		&photo.FileURL,
		&photo.IsMatch,
		&photo.Timestamp,
		&photo.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &photo, nil
}
