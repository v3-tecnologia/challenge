package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/dtos"
)

func TestGeolocationSaveData(t *testing.T) {
	data := dtos.GeolocationDataDto{
		Latitude:  1.0,
		Longitude: 2.0,
	}

	result, err := geolocationService.SaveData(data)

	require.NoError(t, err, "Failed to save geolocation data")
	require.NotEmpty(t, result, "Geolocation should not be empty")
	require.Equal(t, data.Latitude, result.Latitude, "Latitude should match")
	require.Equal(t, data.Longitude, result.Longitude, "Longitude should match")
}
