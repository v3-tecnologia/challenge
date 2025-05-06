package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/dtos"
)

func TestGyroscopeSaveData(t *testing.T) {
	data := dtos.GyroscopeDataDto{
		X: 1.0,
		Y: 2.0,
		Z: 3.0,
	}

	result, err := gyroscopeService.SaveData(data)

	require.NoError(t, err, "Failed to save gyroscope data")
	require.NotEmpty(t, result, "Gyroscope should not be empty")
	require.Equal(t, data.X, result.XValue, "XValue should match")
	require.Equal(t, data.Y, result.YValue, "YValue should match")
	require.Equal(t, data.Z, result.ZValue, "ZValue should match")
}
