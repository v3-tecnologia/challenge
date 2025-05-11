package server

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testServer *FiberServer

func TestMain(m *testing.M) {
	testServer = New()
	testServer.RegisterRoutes()

	code := m.Run()
	os.Exit(code)
}

func TestHealthRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	resp, err := testServer.App.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200 OK")

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "Failed to read response body")

	expectedBody := `{"status":"ok"}`
	assert.JSONEq(t, expectedBody, string(body), "Response body does not match expected JSON")
}

func TestGyroscopeRoute_Body_ShouldNotBeValid(t *testing.T) {
	payload := `{
          "deviceId": "123457",
          "x": 1,
          "y": 1,
          "collectedAt": "2025-05-07T17:40:27.527Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := testServer.App.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expected status code 400 Bad Request")

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "Failed to read response body")

	expectedBody := `{
      "errors": [
        {
          "Error": true,
          "FailedField": "Z",
          "Tag": "required",
          "Value": 0
        }
      ]
    }`
	assert.JSONEq(t, expectedBody, string(body), "Response body does not match expected JSON")
}

func TestGpsRoute_Body_ShouldNotBeValid(t *testing.T) {
	payload := `{
          "deviceId": "123457",
          "latitude": 1,
          "collectedAt": "2025-05-07T17:40:27.527Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := testServer.App.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expected status code 400 Bad Request")
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "Failed to read response body")

	expectedBody := `{
          "errors": [
            {
              "Error": true,
              "FailedField": "Longitude",
              "Tag": "required",
              "Value": null
            }
          ]
        }`

	assert.JSONEq(t, expectedBody, string(body), "Response body does not match expected JSON")
}

func TestPhotoRoute_Body_ShouldNotBeValid(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("deviceId", "123457")
	assert.NoError(t, err)
	err = writer.WriteField("collectedAt", "2025-05-07T17:40:27.527Z")

	err = writer.Close()
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/photo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := testServer.App.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expected status code 400 Bad Request")

	responseBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "Failed to read response body")

	expectedBody := `{
          "errors": "Invalid file"
        }`

	assert.JSONEq(t, expectedBody, string(responseBody), "Response body does not match expected JSON")
}
