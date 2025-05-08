package repository

import (
	"github.com/bielgennaro/v3-challenge-cloud/internal/db"
	"github.com/bielgennaro/v3-challenge-cloud/internal/model"
)

func SaveGyroscope(gyroscope *model.Gyroscope) error {
	conn, err := db.GetConnection()
	if err != nil {
		return err
	}

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
