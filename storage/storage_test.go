package storage

import (
	"challenge-v3/models"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) (*PostgresStorage, *sql.DB) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Log("Aviso: Arquivo .env não encontrado, usando variáveis do sistema.")
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)
	storage, err := NewPostgresStorage(connStr)
	require.NoError(t, err)

	err = storage.InitTables()
	require.NoError(t, err)

	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)

	return storage, db
}

func float64Ptr(f float64) *float64 { return &f }

func TestPostgresStorage_SaveGPS(t *testing.T) {
	storage, db := setupTestDB(t)
	defer db.Close()

	_, err := db.Exec("TRUNCATE TABLE gps RESTART IDENTITY")
	require.NoError(t, err)

	testData := models.GPSData{
		DeviceID:  "test-dev-gps",
		Latitude:  float64Ptr(-10.5),
		Longitude: float64Ptr(-35.5),
		Timestamp: time.Now().UTC().Truncate(time.Second),
	}

	err = storage.SaveGPS(&testData)
	require.NoError(t, err)

	var result models.GPSData
	var lat, lon float64
	err = db.QueryRow("SELECT device_id, latitude, longitude, timestamp FROM gps WHERE device_id = $1", "test-dev-gps").Scan(
		&result.DeviceID, &lat, &lon, &result.Timestamp,
	)
	require.NoError(t, err)
	result.Latitude = &lat
	result.Longitude = &lon

	assert.Equal(t, testData.DeviceID, result.DeviceID)
	assert.InDelta(t, *testData.Latitude, *result.Latitude, 0.001)
	assert.InDelta(t, *testData.Longitude, *result.Longitude, 0.001)
	assert.True(t, testData.Timestamp.Equal(result.Timestamp), "Os timestamps deveriam representar o mesmo momento")
}
