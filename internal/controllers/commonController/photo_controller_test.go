package commonController

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"v3-test/internal/enums"
	"v3-test/internal/models/commonModels"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockPhotoUsecase struct {
	mockUploadFunc func(file *multipart.FileHeader, entity enums.PhotoEntity) (commonModels.PhotoModel, error)
}

func (m *mockPhotoUsecase) UploadPhoto(file *multipart.FileHeader, entity enums.PhotoEntity) (commonModels.PhotoModel, error) {
	return m.mockUploadFunc(file, entity)
}

func setupPhotoRouter(controller PhotoController) *gin.Engine {
	router := gin.Default()
	router.POST("/telemetry/photo", controller.UploadTelemetryPhoto)
	return router
}

func TestUploadTelemetryPhoto_Success(t *testing.T) {
	mockUsecase := &mockPhotoUsecase{
		mockUploadFunc: func(file *multipart.FileHeader, entity enums.PhotoEntity) (commonModels.PhotoModel, error) {
			return commonModels.PhotoModel{
				Entity: entity,
				Url:    "files/photos/test.jpg",
			}, nil
		},
	}

	controller := NewPhotoController(mockUsecase)
	router := setupPhotoRouter(controller)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.jpg")
	part.Write([]byte("fake image content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/telemetry/photo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"photo"`)
}

func TestUploadTelemetryPhoto_NoFile(t *testing.T) {
	mockUsecase := &mockPhotoUsecase{}

	controller := NewPhotoController(mockUsecase)
	router := setupPhotoRouter(controller)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.Close()

	req := httptest.NewRequest("POST", "/telemetry/photo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error"`)
}

func TestUploadTelemetryPhoto_UsecaseError(t *testing.T) {
	mockUsecase := &mockPhotoUsecase{
		mockUploadFunc: func(file *multipart.FileHeader, entity enums.PhotoEntity) (commonModels.PhotoModel, error) {
			return commonModels.PhotoModel{}, errors.New("usecase failed")
		},
	}

	controller := NewPhotoController(mockUsecase)
	router := setupPhotoRouter(controller)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.jpg")
	part.Write([]byte("fake image content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/telemetry/photo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"error"`)
}
