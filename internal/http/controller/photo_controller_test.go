package controller_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestPhotoController_RecognizePhoto(t *testing.T) {
	response := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(response)

	var fileContent = []byte("fake image data")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", "test.jpg")
	require.NoError(t, err)
	_, err = io.Copy(part, bytes.NewReader(fileContent))
	require.NoError(t, err)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/telemetry/photo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	context.Request = req

	photoController.RecognizePhoto(context)
	require.Equal(t, http.StatusOK, response.Code)
}
