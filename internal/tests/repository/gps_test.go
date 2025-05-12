package repository

import (
	"errors"
	"regexp"
	"testing"
	"time"
	"v3/internal/domain"
	gps "v3/internal/repository/gps"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDB(t *testing.T) *gorm.DB {
	// Use in-memory SQLite for tests
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open in-memory SQLite database: %v", err)
	}
	// Migrate the GPS schema
	err = db.AutoMigrate(&domain.GPS{})
	if err != nil {
		t.Fatalf("Failed to migrate GPS schema: %v", err)
	}
	return db
}

func TestGPSRepository_Create(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		input          *domain.GPS
		setupMock      func(mock sqlmock.Sqlmock)
		setupDB        func(db *gorm.DB)
		expectedError  error
		validateResult func(t *testing.T, result *domain.GPS, input *domain.GPS)
		useRealDB      bool
	}{
		{
			name: "Successful GPS creation",
			input: &domain.GPS{
				ID:        "123",
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().UTC(),
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			expectedError: nil,
			validateResult: func(t *testing.T, result *domain.GPS, input *domain.GPS) {
				assert.NotNil(t, result)
				assert.Equal(t, input.ID, result.ID)
				assert.Equal(t, input.DeviceID, result.DeviceID)
				assert.Equal(t, input.Latitude, result.Latitude)
				assert.Equal(t, input.Longitude, result.Longitude)
				assert.WithinDuration(t, input.Timestamp, result.Timestamp, time.Second)
			},
			useRealDB: true,
		},
		{
			name: "Database error",
			input: &domain.GPS{
				ID:        "123",
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().UTC(),
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "gps"`)).
					WithArgs("123", "00:0a:95:9d:68:16", sqlmock.AnyArg(), 40.7128, -74.0060, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			},
			expectedError: errors.New("database error"),
			validateResult: func(t *testing.T, result *domain.GPS, input *domain.GPS) {
				assert.Nil(t, result)
			},
			useRealDB: false,
		},
		{
			name: "Constraint violation - empty DeviceID",
			input: &domain.GPS{
				ID:        "123",
				DeviceID:  "", // Should violate not null constraint
				Timestamp: time.Now().UTC(),
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			expectedError: errors.New("NOT NULL constraint failed"), // SQLite-specific error
			validateResult: func(t *testing.T, result *domain.GPS, input *domain.GPS) {
				assert.Nil(t, result)
			},
			useRealDB: true,
		},
		{
			name: "Duplicate ID violation",
			input: &domain.GPS{
				ID:        "123",
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().UTC(),
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			setupDB: func(db *gorm.DB) {
				// Insert a GPS record with ID "123" to cause a duplicate key violation
				err := db.Create(&domain.GPS{
					ID:        "123",
					DeviceID:  "00:0a:95:9d:68:16",
					Timestamp: time.Now().UTC(),
					Latitude:  40.7128,
					Longitude: -74.0060,
				}).Error
				assert.NoError(t, err)
			},
			expectedError: errors.New("UNIQUE constraint failed"), // SQLite-specific error
			validateResult: func(t *testing.T, result *domain.GPS, input *domain.GPS) {
				assert.Nil(t, result)
			},
			useRealDB: true,
		},
		{
			name: "Constraint violation - zero timestamp",
			input: &domain.GPS{
				ID:        "123",
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Time{}, // Zero timestamp, should violate not null
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			expectedError: errors.New("NOT NULL constraint failed"), // SQLite-specific error
			validateResult: func(t *testing.T, result *domain.GPS, input *domain.GPS) {
				assert.Nil(t, result)
			},
			useRealDB: true,
		},
		{
			name: "Invalid latitude - out of range",
			input: &domain.GPS{
				ID:        "123",
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().UTC(),
				Latitude:  1000, // Invalid latitude (> 90)
				Longitude: -74.0060,
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "gps"`)).
					WithArgs("123", "00:0a:95:9d:68:16", sqlmock.AnyArg(), 1000, -74.0060, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("CHECK constraint failed: latitude"))
				mock.ExpectRollback()
			},
			expectedError: errors.New("CHECK constraint failed"), // Adjust if schema has CHECK
			validateResult: func(t *testing.T, result *domain.GPS, input *domain.GPS) {
				assert.Nil(t, result)
			},
			useRealDB: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var db *gorm.DB
			var repo gps.GPSRepository

			if tt.useRealDB {
				// Use in-memory SQLite for real DB tests
				db = setupDB(t)
				if tt.setupDB != nil {
					tt.setupDB(db)
				}
				repo = gps.NewGPSRepository(db)
			} else {
				// Use sqlmock for error cases
				sqlDB, mock, err := sqlmock.New()
				assert.NoError(t, err)
				defer sqlDB.Close()

				db, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
				assert.NoError(t, err)

				tt.setupMock(mock)
				repo = gps.NewGPSRepository(db)

				// Verify mock expectations
				defer func() {
					assert.NoError(t, mock.ExpectationsWereMet())
				}()
			}

			// Execute the Create method
			result, err := repo.Create(tt.input)

			// Assert error
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			// Validate result
			tt.validateResult(t, result, tt.input)
		})
	}
}
