package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yanvic/challenge/core/entity"
	"github.com/yanvic/challenge/infra/database/dynamo"
	"github.com/yanvic/challenge/internal/handler"
	"github.com/yanvic/challenge/internal/queue"
)

func TestMain(m *testing.M) {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println("⚠️ .env not found, relying on system env")
	// }

	ctx := context.TODO()

	client, err := dynamo.InitDynamoClientTest(ctx)
	if err != nil {
		log.Fatalf("failed to initialize dynamodb client: %v", err)
	}

	dynamo.Client = client

	os.Exit(m.Run())
}

func TestHandlerGyroscope_Success(t *testing.T) {
	x, y, z := 1.0, 2.0, 3.0
	data := entity.Gyroscope{
		X:         &x,
		Y:         &y,
		Z:         &z,
		DeviceID:  "device-gyroscope",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	payload, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.HandlerGyroscope(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Gyroscope data saved")
}

func TestHandlerGyroscope_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", strings.NewReader("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.HandlerGyroscope(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid JSON")
}

func TestHandlerGyroscope_MissingFields(t *testing.T) {
	data := entity.Gyroscope{DeviceID: "dev-1", Timestamp: time.Now().Format(time.RFC3339)}
	payload, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.HandlerGyroscope(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "x is required")
}

func TestHandlerGPS_Success(t *testing.T) {
	lat, lon := -23.5, -46.6
	data := entity.GPS{
		Latitude:  &lat,
		Longitude: &lon,
		DeviceID:  "device-gps",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	payload, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.HandlerGPS(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "GPS data received")
}

func TestHandlerGPS_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", strings.NewReader("{"))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.HandlerGPS(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid JSON")
}

func TestHandlerGPS_MissingFields(t *testing.T) {
	data := entity.GPS{DeviceID: "dev-1"}
	payload, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.HandlerGPS(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandlerGPS_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/telemetry/gps", nil)
	rr := httptest.NewRecorder()

	handler.HandlerGPS(rr, req)
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	assert.Contains(t, rr.Body.String(), "Method not allowed")
}

func TestHandlerPhoto_Success(t *testing.T) {
	handler.PublishImageFunc = func(data []byte) error {
		return nil
	}
	defer func() { handler.PublishImageFunc = queue.PublishImage }()

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	f, err := w.CreateFormFile("image", "test.jpg")
	assert.NoError(t, err)
	_, err = io.Copy(f, strings.NewReader("fakeimagecontent"))
	assert.NoError(t, err)

	w.WriteField("device_id", "photo-device")
	w.WriteField("timestamp", time.Now().Format(time.RFC3339))
	w.Close()

	req := httptest.NewRequest(http.MethodPost, "/telemetry/photo", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())
	rr := httptest.NewRecorder()

	handler.HandlerPhoto(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Image uploaded successfully")
}

func TestHandlerPhoto_MissingImage(t *testing.T) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Não adiciona campo "image"
	w.WriteField("device_id", "photo-device")
	w.WriteField("timestamp", time.Now().Format(time.RFC3339))
	w.Close()

	req := httptest.NewRequest(http.MethodPost, "/telemetry/photo", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())
	rr := httptest.NewRecorder()

	handler.HandlerPhoto(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Image not uploaded")
}

func TestHandlerPhoto_BadMultipartForm(t *testing.T) {
	// Request com corpo inválido que falha no ParseMultipartForm
	req := httptest.NewRequest(http.MethodPost, "/telemetry/photo", strings.NewReader("invalid-data"))
	req.Header.Set("Content-Type", "multipart/form-data; boundary=wrong")

	rr := httptest.NewRecorder()
	handler.HandlerPhoto(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Error processing image")
}

func TestHandlerPhoto_PublishError(t *testing.T) {
	handler.PublishImageFunc = func(data []byte) error {
		return errors.New("publish error")
	}
	defer func() { handler.PublishImageFunc = queue.PublishImage }()

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	f, err := w.CreateFormFile("image", "test.jpg")
	assert.NoError(t, err)
	_, err = io.Copy(f, strings.NewReader("fakeimagecontent"))
	assert.NoError(t, err)

	w.WriteField("device_id", "photo-device")
	w.WriteField("timestamp", time.Now().Format(time.RFC3339))
	w.Close()

	req := httptest.NewRequest(http.MethodPost, "/telemetry/photo", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())
	rr := httptest.NewRecorder()

	handler.HandlerPhoto(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "Error publishing to queue")
}
