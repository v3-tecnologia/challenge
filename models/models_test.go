package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func float64Ptr(f float64) *float64 {
	return &f
}

func TestGyroscopeData_Validate(t *testing.T) {
	baseTime := time.Now()
	validData := func() GyroscopeData {
		return GyroscopeData{
			DeviceID:  "dev-1",
			X:         float64Ptr(1.1),
			Y:         float64Ptr(2.2),
			Z:         float64Ptr(3.3),
			Timestamp: baseTime,
		}
	}

	t.Run("sucesso_payload_completo", func(t *testing.T) {
		data := validData()
		err := data.Validate()
		assert.NoError(t, err)
	})

	testCases := []struct {
		name          string
		mutator       func(*GyroscopeData)
		expectedError string
	}{
		{"falha_sem_device_id", func(d *GyroscopeData) { d.DeviceID = "" }, "campo obrigatório ausente: device_id"},
		{"falha_sem_timestamp", func(d *GyroscopeData) { d.Timestamp = time.Time{} }, "campo obrigatório ausente: timestamp"},
		{"falha_sem_x", func(d *GyroscopeData) { d.X = nil }, "campo obrigatório ausente: x"},
		{"falha_sem_y", func(d *GyroscopeData) { d.Y = nil }, "campo obrigatório ausente: y"},
		{"falha_sem_z", func(d *GyroscopeData) { d.Z = nil }, "campo obrigatório ausente: z"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := validData()
			tc.mutator(&data)
			err := data.Validate()
			assert.Error(t, err)
			assert.Equal(t, tc.expectedError, err.Error())
		})
	}
}

func TestGPSData_Validate(t *testing.T) {
	baseTime := time.Now()
	validData := func() GPSData {
		return GPSData{
			DeviceID:  "dev-2",
			Latitude:  float64Ptr(-8.0),
			Longitude: float64Ptr(-34.0),
			Timestamp: baseTime,
		}
	}

	t.Run("sucesso_payload_completo", func(t *testing.T) {
		data := validData()
		err := data.Validate()
		assert.NoError(t, err)
	})

	testCases := []struct {
		name          string
		mutator       func(*GPSData)
		expectedError string
	}{
		{"falha_sem_device_id", func(d *GPSData) { d.DeviceID = "" }, "campo obrigatório ausente: device_id"},
		{"falha_sem_timestamp", func(d *GPSData) { d.Timestamp = time.Time{} }, "campo obrigatório ausente: timestamp"},
		{"falha_sem_latitude", func(d *GPSData) { d.Latitude = nil }, "campo obrigatório ausente: latitude"},
		{"falha_sem_longitude", func(d *GPSData) { d.Longitude = nil }, "campo obrigatório ausente: longitude"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := validData()
			tc.mutator(&data)
			err := data.Validate()
			assert.Error(t, err)
			assert.Equal(t, tc.expectedError, err.Error())
		})
	}
}

func TestPhotoData_Validate(t *testing.T) {
	baseTime := time.Now()
	validData := func() PhotoData {
		return PhotoData{
			DeviceID:  "dev-3",
			Photo:     "base64stringsimulada",
			Timestamp: baseTime,
		}
	}

	t.Run("sucesso_payload_completo", func(t *testing.T) {
		data := validData()
		err := data.Validate()
		assert.NoError(t, err)
	})

	testCases := []struct {
		name          string
		mutator       func(*PhotoData)
		expectedError string
	}{
		{"falha_sem_device_id", func(d *PhotoData) { d.DeviceID = "" }, "campo obrigatório ausente: device_id"},
		{"falha_sem_timestamp", func(d *PhotoData) { d.Timestamp = time.Time{} }, "campo obrigatório ausente: timestamp"},
		{"falha_sem_photo", func(d *PhotoData) { d.Photo = "" }, "campo obrigatório ausente: photo"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := validData()
			tc.mutator(&data)
			err := data.Validate()
			assert.Error(t, err)
			assert.Equal(t, tc.expectedError, err.Error())
		})
	}
}
