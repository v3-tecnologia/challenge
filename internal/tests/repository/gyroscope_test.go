package repository

import (
	"testing"
	"time"
	"v3/internal/domain"
	"v3/internal/repository/gyroscope"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupDBGyroscope configures an in-memory SQLite database for tests
func setupDBGyroscope(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open in-memory SQLite database: %v", err)
	}

	// Drop the Gyroscope table to ensure a clean state
	err = db.Exec("DROP TABLE IF EXISTS gyroscopes").Error
	if err != nil {
		t.Fatalf("Failed to drop Gyroscope table: %v", err)
	}

	// Migrate the Gyroscope schema
	err = db.AutoMigrate(&domain.Gyroscope{})
	if err != nil {
		t.Fatalf("Failed to migrate Gyroscope schema: %v", err)
	}

	// Debug: Log the table schema
	var schema string
	db.Raw("SELECT sql FROM sqlite_master WHERE type='table' AND name='gyroscopes'").Scan(&schema)
	t.Logf("Gyroscope Table Schema: %s", schema)

	return db
}

func TestGyroscopeRepository_Create(t *testing.T) {
	// Case 1: Successful Gyroscope creation
	t.Run("Successful Gyroscope creation", func(t *testing.T) {
		db := setupDBGyroscope(t)
		repo := gyroscope.NewGyroscopeRepository(db)

		input := &domain.Gyroscope{
			ID:        "123",
			DeviceID:  "00:0a:95:9d:68:16",
			X:         1.5,
			Y:         -2.3,
			Z:         0.8,
			Timestamp: time.Now().UTC(),
		}

		result, err := repo.Create(input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, input.ID, result.ID)
		assert.Equal(t, input.DeviceID, result.DeviceID)
		assert.Equal(t, input.X, result.X)
		assert.Equal(t, input.Y, result.Y)
		assert.Equal(t, input.Z, result.Z)
		assert.WithinDuration(t, input.Timestamp, result.Timestamp, time.Second)
		assert.WithinDuration(t, time.Now().UTC(), result.CreatedAt, time.Second)
	})

	// Case 2: Duplicate ID violation
	t.Run("Duplicate ID violation", func(t *testing.T) {
		db := setupDBGyroscope(t)
		repo := gyroscope.NewGyroscopeRepository(db)

		firstGyro := &domain.Gyroscope{
			ID:        "123",
			DeviceID:  "00:0a:95:9d:68:16",
			X:         1.5,
			Y:         -2.3,
			Z:         0.8,
			Timestamp: time.Now().UTC(),
		}
		_, err := repo.Create(firstGyro)
		assert.NoError(t, err)

		secondGyro := &domain.Gyroscope{
			ID:        "123",
			DeviceID:  "00:0a:95:9d:68:17",
			X:         2.0,
			Y:         -1.8,
			Z:         0.9,
			Timestamp: time.Now().UTC(),
		}
		result, err := repo.Create(secondGyro)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "UNIQUE constraint failed")
		assert.Nil(t, result)
	})

	// Case 3: Empty DeviceID violation
	t.Run("Empty DeviceID violation", func(t *testing.T) {
		db := setupDBGyroscope(t)
		repo := gyroscope.NewGyroscopeRepository(db)

		input := &domain.Gyroscope{
			ID:        "124",
			DeviceID:  "", // Should violate NOT NULL
			X:         1.5,
			Y:         -2.3,
			Z:         0.8,
			Timestamp: time.Now().UTC(),
		}

		result, err := repo.Create(input)

		assert.Error(t, err, "Expected error due to empty DeviceID")
		assert.Contains(t, err.Error(), "NOT NULL constraint failed: gyroscopes.device_id")
		assert.Nil(t, result)
	})

	// Case 4: Zero X violation
	t.Run("Zero X violation", func(t *testing.T) {
		db := setupDBGyroscope(t)
		repo := gyroscope.NewGyroscopeRepository(db)

		input := &domain.Gyroscope{
			ID:        "125",
			DeviceID:  "00:0a:95:9d:68:16",
			X:         0, // Should violate NOT NULL
			Y:         -2.3,
			Z:         0.8,
			Timestamp: time.Now().UTC(),
		}

		result, err := repo.Create(input)

		assert.Error(t, err, "Expected error due to zero X")
		assert.Contains(t, err.Error(), "NOT NULL constraint failed: gyroscopes.x")
		assert.Nil(t, result)
	})

	// Case 5: Zero Timestamp violation
	t.Run("Zero Timestamp violation", func(t *testing.T) {
		db := setupDBGyroscope(t)
		repo := gyroscope.NewGyroscopeRepository(db)

		input := &domain.Gyroscope{
			ID:        "126",
			DeviceID:  "00:0a:95:9d:68:16",
			X:         1.5,
			Y:         -2.3,
			Z:         0.8,
			Timestamp: time.Time{}, // Zero timestamp
		}

		result, err := repo.Create(input)

		assert.Error(t, err, "Expected error due to zero Timestamp")
		assert.Contains(t, err.Error(), "NOT NULL constraint failed: gyroscopes.timestamp")
		assert.Nil(t, result)
	})
}
