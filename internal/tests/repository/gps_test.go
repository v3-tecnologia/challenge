package repository

import (
	"testing"
	"time"
	"v3/internal/domain"
	"v3/internal/repository/gps"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupDB configures an in-memory SQLite database for tests
func setupDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open in-memory SQLite database: %v", err)
	}

	// Drop the GPS table to ensure a clean state
	err = db.Exec("DROP TABLE IF EXISTS gps").Error
	if err != nil {
		t.Fatalf("Failed to drop GPS table: %v", err)
	}

	// Migrate the GPS schema
	err = db.AutoMigrate(&domain.GPS{})
	if err != nil {
		t.Fatalf("Failed to migrate GPS schema: %v", err)
	}

	// Debug: Log the table schema
	var schema string
	db.Raw("SELECT sql FROM sqlite_master WHERE type='table' AND name='gps'").Scan(&schema)
	t.Logf("GPS Table Schema: %s", schema)

	return db
}

func TestGPSRepository_Create(t *testing.T) {
	// Case 1: Successful GPS creation
	t.Run("Successful GPS creation", func(t *testing.T) {
		db := setupDB(t)
		repo := gps.NewGPSRepository(db)

		input := &domain.GPS{
			ID:        "123",
			DeviceID:  "00:0a:95:9d:68:16",
			Timestamp: time.Now().UTC(),
			Latitude:  40.7128,
			Longitude: -74.0060,
		}

		result, err := repo.Create(input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, input.ID, result.ID)
		assert.Equal(t, input.DeviceID, result.DeviceID)
		assert.Equal(t, input.Latitude, result.Latitude)
		assert.Equal(t, input.Longitude, result.Longitude)
		assert.WithinDuration(t, input.Timestamp, result.Timestamp, time.Second)
		assert.WithinDuration(t, time.Now().UTC(), result.CreatedAt, time.Second)
	})

	// Case 2: Duplicate ID violation
	t.Run("Duplicate ID violation", func(t *testing.T) {
		db := setupDB(t)
		repo := gps.NewGPSRepository(db)

		firstGPS := &domain.GPS{
			ID:        "123",
			DeviceID:  "00:0a:95:9d:68:16",
			Timestamp: time.Now().UTC(),
			Latitude:  40.7128,
			Longitude: -74.0060,
		}
		_, err := repo.Create(firstGPS)
		assert.NoError(t, err)

		secondGPS := &domain.GPS{
			ID:        "123",
			DeviceID:  "00:0a:95:9d:68:17",
			Timestamp: time.Now().UTC(),
			Latitude:  41.7128,
			Longitude: -75.0060,
		}
		result, err := repo.Create(secondGPS)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "UNIQUE constraint failed")
		assert.Nil(t, result)
	})

	// Case 3: Empty DeviceID violation
	t.Run("Empty DeviceID violation", func(t *testing.T) {
		db := setupDB(t)
		repo := gps.NewGPSRepository(db)

		input := &domain.GPS{
			ID:        "124",
			DeviceID:  "", // Should violate NOT NULL
			Timestamp: time.Now().UTC(),
			Latitude:  40.7128,
			Longitude: -74.0060,
		}

		result, err := repo.Create(input)

		assert.Error(t, err, "Expected error due to empty DeviceID, but got nil")
		if err != nil {
			assert.Contains(t, err.Error(), "NOT NULL constraint failed")
		}
		assert.Nil(t, result)
	})

	// Case 4: Zero Timestamp violation
	t.Run("Zero Timestamp violation", func(t *testing.T) {
		db := setupDB(t)
		repo := gps.NewGPSRepository(db)

		input := &domain.GPS{
			ID:        "125",
			DeviceID:  "00:0a:95:9d:68:16",
			Timestamp: time.Time{}, // Zero timestamp
			Latitude:  40.7128,
			Longitude: -74.0060,
		}

		result, err := repo.Create(input)

		assert.Error(t, err)
		if err != nil {
			assert.Contains(t, err.Error(), "NOT NULL constraint failed")
		}
		assert.Nil(t, result)
	})

	// Case 5: Missing Latitude violation
	t.Run("Missing Latitude violation", func(t *testing.T) {
		db := setupDB(t)
		repo := gps.NewGPSRepository(db)

		input := &domain.GPS{
			ID:        "126",
			DeviceID:  "00:0a:95:9d:68:16",
			Timestamp: time.Now().UTC(),
			Latitude:  0, // Should violate NOT NULL
			Longitude: -74.0060,
		}

		result, err := repo.Create(input)

		assert.Error(t, err)
		if err != nil {
			assert.Contains(t, err.Error(), "NOT NULL constraint failed")
		}
		assert.Nil(t, result)
	})
}
