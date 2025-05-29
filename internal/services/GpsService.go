package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/KaiRibeiro/challenge/internal/logs"
	"github.com/KaiRibeiro/challenge/internal/models"
)

type GpsService interface {
	AddGps(gps models.GpsModel) error
}

type GpsDBService struct {
	DB *sql.DB
}

func NewGPSDBService(dbConn *sql.DB) *GpsDBService {
	return &GpsDBService{DB: dbConn}
}

func (s *GpsDBService) AddGps(gps models.GpsModel) error {
	logs.Logger.Info("adding gps data to database",
		"latitude", gps.Latitude,
		"longitude", gps.Longitude,
		"mac", gps.MAC,
		"timestamp", gps.Timestamp,
	)

	timestamp := time.UnixMilli(gps.Timestamp)

	query := `INSERT INTO gps (latitude, longitude, mac, timestamp)
    VALUES ($1, $2, $3, $4)`

	_, err := s.DB.Exec(query, gps.Latitude, gps.Longitude, gps.MAC, timestamp)

	if err != nil {
		wrappedErr := fmt.Errorf("failed to insert gps data into database: %w", custom_errors.NewDBError(err, http.StatusInternalServerError))
		logs.Logger.Error("failed to add gps data",
			"error", wrappedErr,
		)
		return wrappedErr
	}

	return err
}
