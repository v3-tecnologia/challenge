package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/infra/database"
	"github.com/wellmtx/challenge/internal/infra/entities"
	"github.com/wellmtx/challenge/internal/infra/repositories"
)

func TestGyroscopeCreate(t *testing.T) {
	// Initialize the database connection
	db := database.NewDatabase(
		"test",
		"test",
		"test",
		"test",
		true,
	)
	err := db.Connect()
	require.NoError(t, err, "Failed to connect to the database")
	defer db.Close()

	// Initialize the Gyroscope repository
	repo := repositories.NewGyroscopeRepository(db)

	// Create a new Gyroscope entity
	gyroscope := entities.GyroscopeEntity{
		XValue: 1.0,
		YValue: 2.0,
		ZValue: 3.0,
	}

	// Call the Create method
	result, err := repo.Create(gyroscope)
	require.NoError(t, err, "Failed to create gyroscope")

	// Verify that the gyroscope was created in the database
	var createdGyroscope entities.GyroscopeEntity
	err = db.DB.First(&createdGyroscope, result.ID).Error
	require.NoError(t, err, "Failed to find created gyroscope")
	require.Equal(t, result.ID, createdGyroscope.ID)
	require.Equal(t, result.XValue, createdGyroscope.XValue)
	require.Equal(t, result.YValue, createdGyroscope.YValue)
	require.Equal(t, result.ZValue, createdGyroscope.ZValue)
}
