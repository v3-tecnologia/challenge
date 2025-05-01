package services

import (
	"github.com/KaiRibeiro/challenge/internal/db"
	"github.com/KaiRibeiro/challenge/internal/models"
	"time"
)

func AddGyroscope(gyroscope models.GyroscopeModel) error {
	timestamp := time.UnixMilli(gyroscope.Timestamp)

	query := `INSERT INTO gyroscope (x, y, z, mac, timestamp)
    VALUES ($1, $2, $3, $4, $5)`

	_, err := db.DB.Exec(query, gyroscope.X, gyroscope.Y, gyroscope.Z, gyroscope.MAC, timestamp)

	return err
}
