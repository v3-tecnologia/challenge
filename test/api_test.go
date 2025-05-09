package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"v3-backend-challenge/src/dto"
)

func TestHandleGyroscope(t *testing.T) {
	db := SetupTestDB()
	router := SetupRouter(db)

	payload := dto.Gyroscope{
		AxisX: 1.2,
		AxisY: -0.8,
		AxisZ: 0.4,
		BaseTelemetry: dto.BaseTelemetry{
			DateTimeCollected: time.Now(),
			MacAddr:           "00:11:22:33:44:55",
		},
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestHandleGPS(t *testing.T) {
	db := SetupTestDB()
	router := SetupRouter(db)

	payload := dto.GPS{
		BaseTelemetry: dto.BaseTelemetry{
			DateTimeCollected: time.Now(),
			MacAddr:           "00:11:22:33:44:55",
		},
		Latitude:  30,
		Longitude: 4,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/telemetry/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestHandlePhoto(t *testing.T) {
	db := SetupTestDB()
	router := SetupRouter(db)

	var img bytes.Buffer
	writer := multipart.NewWriter(&img)

	_ = writer.WriteField("datetime_collected", time.Now().Format(time.RFC3339))
	_ = writer.WriteField("mac_addr", "00:11:22:33:44:55")

	fileWriter, err := writer.CreateFormFile("img", "test.jpg")
	require.NoError(t, err)

	imageContent := []byte("teste")
	_, err = fileWriter.Write(imageContent)
	require.NoError(t, err)

	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "/telemetry/photo", &img)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestHandlePhoto_Errors(t *testing.T) {
	db := SetupTestDB()
	router := SetupRouter(db)

	tests := []struct {
		name         string
		fields       map[string]string
		includeImage bool
		expectedCode int
		expectedMsg  string
	}{
		{
			name:         "datetime_collected não fornecido",
			fields:       map[string]string{"mac_addr": "00:11:22:33:44:55"},
			includeImage: true,
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "datetime_collected não fornecido",
		},
		{
			name:         "datetime_collected inválido",
			fields:       map[string]string{"datetime_collected": "invalido", "mac_addr": "00:11:22:33:44:55"},
			includeImage: true,
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "datetime_collected inválido",
		},
		{
			name:         "Mac address não fornecido",
			fields:       map[string]string{"datetime_collected": time.Now().Format(time.RFC3339)},
			includeImage: true,
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "Mac address não fornecido",
		},
		{
			name:         "Foto não fornecida",
			fields:       map[string]string{"datetime_collected": time.Now().Format(time.RFC3339), "mac_addr": "00:11:22:33:44:55"},
			includeImage: false,
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "Foto não fornecida",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var body bytes.Buffer
			writer := multipart.NewWriter(&body)

			for k, v := range tc.fields {
				_ = writer.WriteField(k, v)
			}

			if tc.includeImage {
				part, _ := writer.CreateFormFile("img", "test.jpg")
				part.Write([]byte("img"))
			}

			writer.Close()

			req, _ := http.NewRequest("POST", "/telemetry/photo", &body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tc.expectedCode, resp.Code)
			assert.Contains(t, resp.Body.String(), tc.expectedMsg)
		})
	}
}

func TestHandleGps_InvalidPayload(t *testing.T) {
	db := SetupTestDB()
	router := SetupRouter(db)

	req, _ := http.NewRequest("POST", "/telemetry/gps", nil)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "Payload não está conforme o esperado")
}

func TestHandleGyroscope_InvalidPayload(t *testing.T) {
	db := SetupTestDB()
	router := SetupRouter(db)

	req, _ := http.NewRequest("POST", "/telemetry/gyroscope", nil)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "Payload não está conforme o esperado")
}
