package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/infra/database"
	"github.com/wellmtx/challenge/internal/infra/entities"
	"github.com/wellmtx/challenge/internal/infra/repositories"
)

func TestPhotoRepository_Create(t *testing.T) {
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

	repo := repositories.NewPhotoRepository(db)

	photo := entities.PhotoEntity{
		FilePath: "test/path/to/photo.jpg",
	}

	createdPhoto, err := repo.Create(photo)
	require.NoError(t, err, "Failed to create photo")
	require.Equal(t, photo.FilePath, createdPhoto.FilePath, "File paths do not match")
}

func TestPhotoRepository_GetAll(t *testing.T) {
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

	repo := repositories.NewPhotoRepository(db)

	photo1 := entities.PhotoEntity{
		FilePath: "test/path/to/photo1.jpg",
	}
	photo2 := entities.PhotoEntity{
		FilePath: "test/path/to/photo2.jpg",
	}

	_, err = repo.Create(photo1)
	require.NoError(t, err, "Failed to create photo1")
	_, err = repo.Create(photo2)
	require.NoError(t, err, "Failed to create photo2")

	photos, err := repo.GetAll()
	require.NoError(t, err, "Failed to get all photos")
	require.Len(t, photos, 2, "Expected 2 photos in the database")
	require.Equal(t, photo1.FilePath, photos[0].FilePath, "File paths do not match")
	require.Equal(t, photo2.FilePath, photos[1].FilePath, "File paths do not match")
}

func TestPhotoRepository_ListByMacAddress(t *testing.T) {
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

	repo := repositories.NewPhotoRepository(db)

	macAddress := "00:00:00:00:00:01"
	photo1 := entities.PhotoEntity{
		FilePath: "test/path/to/photo1.jpg",
		BaseEntity: entities.BaseEntity{
			MacAddress: macAddress,
		},
	}
	photo2 := entities.PhotoEntity{
		FilePath: "test/path/to/photo2.jpg",
		BaseEntity: entities.BaseEntity{
			MacAddress: macAddress,
		},
	}
	photo3 := entities.PhotoEntity{
		FilePath: "test/path/to/photo3.jpg",
		BaseEntity: entities.BaseEntity{
			MacAddress: "00:00:00:00:00:02",
		},
	}

	_, err = repo.Create(photo1)
	require.NoError(t, err, "Failed to create photo1")
	_, err = repo.Create(photo2)
	require.NoError(t, err, "Failed to create photo2")
	_, err = repo.Create(photo3)
	require.NoError(t, err, "Failed to create photo3")

	photos, err := repo.ListByMacAddress(macAddress)
	require.NoError(t, err, "Failed to list photos by MAC address")
	require.Len(t, photos, 2, "Expected 2 photos for the given MAC address")
	require.Equal(t, photo1.FilePath, photos[0].FilePath, "File paths do not match")
	require.Equal(t, photo2.FilePath, photos[1].FilePath, "File paths do not match")
}
