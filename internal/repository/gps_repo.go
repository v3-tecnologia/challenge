package repository

import (
	"database/sql"

	"github.com/bielgennaro/v3-challenge-cloud/internal/db"
	"github.com/bielgennaro/v3-challenge-cloud/internal/model"
	"github.com/google/uuid"
)

func SaveGPS(gps *model.GPS) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}

	// Insertion Query
	query := `
        INSERT INTO gps_data (id, mac_address, latitude, longitude, timestamp, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `

	_, err = conn.Exec(
		query,
		gps.ID,
		gps.MacAddress,
		gps.Latitude,
		gps.Longitude,
		gps.Timestamp,
		gps.CreatedAt,
	)

	return err
}

func GetGPSByID(id uuid.UUID) (*model.GPS, error) {
	conn, err := db.GetConnection()
	if err != nil {
		return nil, err
	}

	query := `
        SELECT id, mac_address, latitude, longitude, timestamp, created_at
        FROM gps_data
        WHERE id = $1
    `

	var gps model.GPS
	err = conn.QueryRow(query, id).Scan(
		&gps.ID,
		&gps.MacAddress,
		&gps.Latitude,
		&gps.Longitude,
		&gps.Timestamp,
		&gps.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &gps, nil
}
