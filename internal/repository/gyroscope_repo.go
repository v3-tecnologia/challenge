package repository

import (
	"database/sql"

	"github.com/bielgennaro/v3-challenge-cloud/internal/db"
	"github.com/bielgennaro/v3-challenge-cloud/internal/model"
	"github.com/google/uuid"
)

func SaveGyroscope(gyroscope *model.Gyroscope) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}

	// Insertion Query
	query := `
        INSERT INTO gyroscope_data (id, mac_address, axis_x, axis_y, axis_z, timestamp, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

	_, err = conn.Exec(
		query,
		gyroscope.ID,
		gyroscope.MacAddress,
		gyroscope.AxisX,
		gyroscope.AxisY,
		gyroscope.AxisZ,
		gyroscope.Timestamp,
		gyroscope.CreatedAt,
	)

	return err
}

func GetGyroscopeByID(id uuid.UUID) (*model.Gyroscope, error) {
	conn, err := db.GetConnection()
	if err != nil {
		return nil, err
	}

	query := `
        SELECT id, mac_address, axis_x, axis_y, axis_z, timestamp, created_at
        FROM gyroscope_data
        WHERE id = $1
    `

	var gyroscope model.Gyroscope
	err = conn.QueryRow(query, id).Scan(
		&gyroscope.ID,
		&gyroscope.MacAddress,
		&gyroscope.AxisX,
		&gyroscope.AxisY,
		&gyroscope.AxisZ,
		&gyroscope.Timestamp,
		&gyroscope.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &gyroscope, nil
}
