package integration

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTelemetryPhoto(t *testing.T) {
	ClearDatabase(t)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "fake-photo.jpg")
	assert.NoError(t, err)

	randomContent := []byte("this is a fake image content for test")
	_, err = io.Copy(part, bytes.NewReader(randomContent))
	assert.NoError(t, err)

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/telemetry/photo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"photo":`)
}
