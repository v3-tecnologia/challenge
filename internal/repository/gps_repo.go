package repository

import (
	"fmt"
	"log"

	"github.com/bielgennaro/v3-challenge-cloud/internal/db"
	"github.com/bielgennaro/v3-challenge-cloud/internal/model"
)

func SaveGPS(gps *model.GPS) error {
	conn, err := db.GetConnection()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return fmt.Errorf("database connection error: %w", err)
	}
	defer conn.Close()

	// Log dos valores que estão sendo inseridos
	log.Printf("Saving GPS data: ID=%s, MacAddress=%s, Lat=%.6f, Lng=%.6f, Timestamp=%v",
		gps.ID, gps.MacAddress, *gps.Latitude, *gps.Longitude, gps.Timestamp)

	// Insertion Query
	query := `
        INSERT INTO gps (id, mac_address, latitude, longitude, timestamp, created_at)
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

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return fmt.Errorf("database insertion error: %w", err)
	}

	log.Printf("GPS data saved successfully with ID: %s", gps.ID)
	return nil
}
