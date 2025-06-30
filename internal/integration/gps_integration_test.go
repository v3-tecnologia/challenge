package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	controller "challenge-cloud/internal/controllers"
	"challenge-cloud/internal/models"
	repository "challenge-cloud/internal/repositories/gorm"
	router "challenge-cloud/internal/router"
	service "challenge-cloud/internal/services"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupServerGPS(t *testing.T) (http.Handler, *gorm.DB) {

	dsn := "root:@tcp(127.0.0.1:3306)/telemetry_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "Falha ao conectar ao banco MySQL")

	err = db.Exec("DROP TABLE IF EXISTS gps").Error
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.GPS{})
	assert.NoError(t, err)

	repo := repository.NewGPSRepository(db)
	svc := service.NewGPSService(repo)
	ctrl := controller.NewGPSController(svc)

	controllers := router.Controllers{GPS: ctrl}
	handler := router.LoadRouter(controllers)

	return handler, db
}

func TestIntegration_CreateGPS(t *testing.T) {
	server, db := setupServerGPS(t)

	now := time.Now().UTC().Truncate(time.Second)
	payload := models.GPS{
		Longitude: 7.8,
		Latitude:  13.5,
		MAC:       "AA:BB:CC:DD:EE:FF",
		Timestamp: now,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	fmt.Println(bytes.NewBuffer(body))

	assert.Equal(t, http.StatusCreated, rec.Code)

	var respBody models.GPS
	err := json.Unmarshal(rec.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.NotZero(t, respBody.ID)
	assert.Equal(t, payload.MAC, respBody.MAC)
	assert.Equal(t, payload.Latitude, respBody.Latitude)
	assert.Equal(t, payload.Longitude, respBody.Longitude)
	assert.True(t, respBody.Timestamp.Equal(now))

	var saved models.GPS
	err = db.First(&saved, respBody.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, respBody.ID, saved.ID)

	assert.InDelta(t, payload.Latitude, saved.Latitude, 1e-6)
	assert.InDelta(t, payload.Longitude, saved.Longitude, 1e-6) // ver se ta vlr proxim ja q é float
	assert.Equal(t, payload.MAC, saved.MAC)
}
