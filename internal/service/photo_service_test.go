package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/dtos"
)

func TestPhotoService_RecognizePhoto(t *testing.T) {
	// Recognize the photo
	recognizedPhoto, err := photoService.RecognizePhoto(dtos.SavePhotoDTO{
		Image:    []byte("test"),
		FilePath: "test.jpg",
		BaseDTO: dtos.BaseDTO{
			MacAddress: "00:00:00:00:00:01",
			Timestamp:  1234567890,
		},
	})
	require.NoError(t, err)

	// Check if the recognized photo is the same as the created one
	require.NotNil(t, recognizedPhoto)
	require.NotEmpty(t, recognizedPhoto.ID)
	require.Equal(t, recognizedPhoto.FilePath, "test.jpg")
}
