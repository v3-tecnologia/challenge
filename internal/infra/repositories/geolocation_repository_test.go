package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/infra/database"
	"github.com/wellmtx/challenge/internal/infra/entities"
	"github.com/wellmtx/challenge/internal/infra/repositories"
)

func TestGeolocationCreate(t *testing.T) {
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

	// Initialize the Geolocation repository
	repo := repositories.NewGeolocationRepository(db)

	// Create a new Geolocation entity
	geolocation := entities.GeolocationEntity{
		Latitude:  1.0,
		Longitude: 2.0,
	}

	// Call the Create method
	result, err := repo.Create(geolocation)
	require.NoError(t, err, "Failed to create geolocation")

	// Verify that the geolocation was created in the database
	var createdGeolocation entities.GeolocationEntity
	err = db.DB.First(&createdGeolocation, result.ID).Error
	require.NoError(t, err, "Failed to find created geolocation")
	require.Equal(t, result.ID, createdGeolocation.ID)
	require.Equal(t, result.Latitude, createdGeolocation.Latitude)
	require.Equal(t, result.Longitude, createdGeolocation.Longitude)
}
