package repository

import (
	"testing"
	"time"
	"v3/internal/domain"
	"v3/internal/repository/photo"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupDBPhoto configures an in-memory SQLite database for tests
func setupDBPhoto(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open in-memory SQLite database: %v", err)
	}

	// Drop the Photo table to ensure a clean state
	err = db.Exec("DROP TABLE IF EXISTS photos").Error
	if err != nil {
		t.Fatalf("Failed to drop Photo table: %v", err)
	}

	// Migrate the Photo schema
	err = db.AutoMigrate(&domain.Photo{})
	if err != nil {
		t.Fatalf("Failed to migrate Photo schema: %v", err)
	}

	// Debug: Log the table schema
	var schema string
	db.Raw("SELECT sql FROM sqlite_master WHERE type='table' AND name='photos'").Scan(&schema)
	t.Logf("Photo Table Schema: %s", schema)

	return db
}

func TestPhotoRepository_Create(t *testing.T) {
	// Case 1: Successful Photo creation
	t.Run("Successful Photo creation", func(t *testing.T) {
		db := setupDBPhoto(t)
		repo := photo.NewPhotoRepository(db)

		input := &domain.Photo{
			ID:         "123",
			DeviceID:   "00:0a:95:9d:68:16",
			FilePath:   "/photos/123.jpg",
			Timestamp:  time.Now().UTC(),
			Recognized: false,
		}

		result, err := repo.Create(input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, input.ID, result.ID)
		assert.Equal(t, input.DeviceID, result.DeviceID)
		assert.Equal(t, input.FilePath, result.FilePath)
		assert.Equal(t, input.Recognized, result.Recognized)
		assert.WithinDuration(t, input.Timestamp, result.Timestamp, time.Second)
		assert.WithinDuration(t, time.Now().UTC(), result.CreatedAt, time.Second)
	})

	// Case 2: Duplicate ID violation
	t.Run("Duplicate ID violation", func(t *testing.T) {
		db := setupDBPhoto(t)
		repo := photo.NewPhotoRepository(db)

		firstPhoto := &domain.Photo{
			ID:         "123",
			DeviceID:   "00:0a:95:9d:68:16",
			FilePath:   "/photos/123.jpg",
			Timestamp:  time.Now().UTC(),
			Recognized: false,
		}
		_, err := repo.Create(firstPhoto)
		assert.NoError(t, err)

		secondPhoto := &domain.Photo{
			ID:         "123",
			DeviceID:   "00:0a:95:9d:68:17",
			FilePath:   "/photos/124.jpg",
			Timestamp:  time.Now().UTC(),
			Recognized: true,
		}
		result, err := repo.Create(secondPhoto)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "UNIQUE constraint failed")
		assert.Nil(t, result)
	})

	// Case 3: Empty DeviceID violation
	t.Run("Empty DeviceID violation", func(t *testing.T) {
		db := setupDBPhoto(t)
		repo := photo.NewPhotoRepository(db)

		input := &domain.Photo{
			ID:         "124",
			DeviceID:   "", // Should violate NOT NULL
			FilePath:   "/photos/124.jpg",
			Timestamp:  time.Now().UTC(),
			Recognized: false,
		}

		result, err := repo.Create(input)

		assert.Error(t, err, "Expected error due to empty DeviceID")
		assert.Contains(t, err.Error(), "NOT NULL constraint failed: photos.device_id")
		assert.Nil(t, result)
	})

	// Case 4: Empty FilePath violation
	t.Run("Empty FilePath violation", func(t *testing.T) {
		db := setupDBPhoto(t)
		repo := photo.NewPhotoRepository(db)

		input := &domain.Photo{
			ID:         "125",
			DeviceID:   "00:0a:95:9d:68:16",
			FilePath:   "", // Should violate NOT NULL
			Timestamp:  time.Now().UTC(),
			Recognized: false,
		}

		result, err := repo.Create(input)

		assert.Error(t, err, "Expected error due to empty FilePath")
		assert.Contains(t, err.Error(), "NOT NULL constraint failed: photos.file_path")
		assert.Nil(t, result)
	})

	// Case 5: Zero Timestamp violation
	t.Run("Zero Timestamp violation", func(t *testing.T) {
		db := setupDBPhoto(t)
		repo := photo.NewPhotoRepository(db)

		input := &domain.Photo{
			ID:         "126",
			DeviceID:   "00:0a:95:9d:68:16",
			FilePath:   "/photos/126.jpg",
			Timestamp:  time.Time{}, // Zero timestamp
			Recognized: false,
		}

		result, err := repo.Create(input)

		assert.Error(t, err, "Expected error due to zero Timestamp")
		assert.Contains(t, err.Error(), "NOT NULL constraint failed: photos.timestamp")
		assert.Nil(t, result)
	})
}
