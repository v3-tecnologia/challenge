package services

import (
	"github.com/KaiRibeiro/challenge/internal/db"
	"github.com/KaiRibeiro/challenge/internal/models"
	"time"
)

func AddGps(gps models.GpsModel) error {
	timestamp := time.UnixMilli(gps.Timestamp)

	query := `INSERT INTO gps (latitude, longitude, mac, timestamp)
    VALUES ($1, $2, $3, $4)`

	_, err := db.DB.Exec(query, gps.Latitude, gps.Longitude, gps.MAC, timestamp)

	return err
}
