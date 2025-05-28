package routes_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockGyroscopeService struct {
	ErrToReturn error
}

func (m *MockGyroscopeService) AddGyroscope(gyroscope models.GyroscopeModel) error {
	return m.ErrToReturn
}

func TestSaveGyroscope_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockGyroscopeService{}
	handler := routes.NewGyroscopeHandler(mockService)

	router := gin.Default()
	router.POST("/gyroscope", handler.SaveGyroscope)

	payload := models.GyroscopeModel{
		X:         12.34,
		Y:         56.78,
		Z:         90.12,
		MAC:       "8C:16:45:8D:F3:7B",
		Timestamp: 1746110367207,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Gyroscope Saved Successfully")
}

func TestSaveGyroscope_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockGyroscopeService{}
	handler := routes.NewGyroscopeHandler(mockService)

	router := gin.Default()
	router.POST("/gyroscope", handler.SaveGyroscope)

	req := httptest.NewRequest(http.MethodPost, "/gyroscope", bytes.NewBuffer([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestSaveGyroscope_ValidationErrors(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockGyroscopeService{}
	handler := routes.NewGyroscopeHandler(mockService)

	router := gin.Default()
	router.POST("/gyroscope", handler.SaveGyroscope)

	req := httptest.NewRequest(http.MethodPost, "/gyroscope", bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "errors")
}

func TestSaveGyroscope_DBError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockGyroscopeService{
		ErrToReturn: custom_errors.NewDBError(errors.New("DB failure"), 500),
	}
	handler := routes.NewGyroscopeHandler(mockService)

	router := gin.Default()
	router.POST("/gyroscope", handler.SaveGyroscope)

	payload := models.GyroscopeModel{
		X:         12.34,
		Y:         56.78,
		Z:         90.12,
		MAC:       "8C:16:45:8D:F3:7B",
		Timestamp: 1746110367207,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Contains(t, w.Body.String(), "A database error occurred")
}

func TestSaveGyroscope_UnexpectedError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockGyroscopeService{
		ErrToReturn: errors.New("something went wrong"),
	}
	handler := routes.NewGyroscopeHandler(mockService)

	router := gin.Default()
	router.POST("/gyroscope", handler.SaveGyroscope)

	payload := models.GyroscopeModel{
		X:         12.34,
		Y:         56.78,
		Z:         90.12,
		MAC:       "8C:16:45:8D:F3:7B",
		Timestamp: 1746110367207,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "An unexpected server error occurred")
}
