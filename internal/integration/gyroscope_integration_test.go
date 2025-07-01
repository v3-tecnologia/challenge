package integration

import (
	"bytes"
	"encoding/json"
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

func setupServerGyro(t *testing.T) (http.Handler, *gorm.DB) {

	dsn := "root:root@tcp(127.0.0.1:3306)/telemetry_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "Falha ao conectar ao banco MySQL")

	err = db.Exec("DROP TABLE IF EXISTS gyroscopes").Error
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Gyroscope{})
	assert.NoError(t, err)

	repo := repository.NewGyroscopeRepository(db)
	svc := service.NewGyroscopeService(repo)
	ctrl := controller.NewGyroscopeController(svc)

	controllers := router.Controllers{Gyro: ctrl}
	handler := router.LoadRouter(controllers)

	return handler, db
}

func TestIntegration_CreateGyroscope(t *testing.T) {
	server, db := setupServerGyro(t)

	now := time.Now().UTC().Truncate(time.Second)
	payload := models.Gyroscope{
		X:         1.23,
		Y:         4.56,
		Z:         7.89,
		MAC:       "AA:BB:CC:DD:EE:FF",
		Timestamp: now,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	var respBody models.Gyroscope
	err := json.Unmarshal(rec.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.NotZero(t, respBody.ID)
	assert.Equal(t, payload.MAC, respBody.MAC)
	assert.Equal(t, payload.X, respBody.X)
	assert.Equal(t, payload.Y, respBody.Y)
	assert.Equal(t, payload.Z, respBody.Z)
	assert.True(t, respBody.Timestamp.Equal(now))

	var saved models.Gyroscope
	err = db.First(&saved, respBody.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, respBody.ID, saved.ID)
	assert.InDelta(t, payload.X, saved.X, 1e-6)
	assert.InDelta(t, payload.Y, saved.Y, 1e-6)
	assert.InDelta(t, payload.Z, saved.Z, 1e-6)
	assert.Equal(t, payload.MAC, saved.MAC)
}
