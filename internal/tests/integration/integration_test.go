package integration

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    r := gin.Default()

    // Mock handlers
    r.POST("/telemetry/gyroscope", func(c *gin.Context) {
        var payload map[string]float64
        if err := c.ShouldBindJSON(&payload); err != nil || len(payload) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty payload"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "success"})
    })

    r.POST("/telemetry/gps", func(c *gin.Context) {
        var payload map[string]float64
        if err := c.ShouldBindJSON(&payload); err != nil || len(payload) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty payload"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "success"})
    })

    r.POST("/telemetry/photo", func(c *gin.Context) {
        var payload map[string]string
        if err := c.ShouldBindJSON(&payload); err != nil || len(payload) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty payload"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "success"})
    })

    return r
}

func TestGyroscopeEndpoint(t *testing.T) {
    r := setupRouter()

    // Test with valid data
    validBody := `{"x": 1.0, "y": 2.0, "z": 3.0}`
    req, _ := http.NewRequest("POST", "/telemetry/gyroscope", bytes.NewBufferString(validBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)

    // Test with empty data
    emptyBody := `{}`
    req, _ = http.NewRequest("POST", "/telemetry/gyroscope", bytes.NewBufferString(emptyBody))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGPSEndpoint(t *testing.T) {
    r := setupRouter()

    // Test with valid data
    validBody := `{"latitude": 12.34, "longitude": 56.78}`
    req, _ := http.NewRequest("POST", "/telemetry/gps", bytes.NewBufferString(validBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)

    // Test with empty data
    emptyBody := `{}`
    req, _ = http.NewRequest("POST", "/telemetry/gps", bytes.NewBufferString(emptyBody))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPhotoEndpoint(t *testing.T) {
    r := setupRouter()

    // Test with valid data
    validBody := `{"image": "base64encodedstring"}`
    req, _ := http.NewRequest("POST", "/telemetry/photo", bytes.NewBufferString(validBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)

    // Test with empty data
    emptyBody := `{}`
    req, _ = http.NewRequest("POST", "/telemetry/photo", bytes.NewBufferString(emptyBody))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusBadRequest, w.Code)
}