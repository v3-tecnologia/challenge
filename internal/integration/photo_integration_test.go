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

func setupServerPhoto(t *testing.T) (http.Handler, *gorm.DB) {

	dsn := "root:@tcp(127.0.0.1:3306)/telemetry_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "Falha ao conectar ao banco MySQL")

	err = db.Exec("DROP TABLE IF EXISTS gyroscopes").Error
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Photo{})
	assert.NoError(t, err)

	repo := repository.NewPhotoRepository(db)
	svc := service.NewPhotoService(repo)
	ctrl := controller.NewPhotoController(svc)

	controllers := router.Controllers{Photo: ctrl}
	handler := router.LoadRouter(controllers)

	return handler, db
}

func TestIntegration_CreatePhoto(t *testing.T) {
	server, db := setupServerPhoto(t)

	now := time.Now().UTC().Truncate(time.Second)
	payload := models.Photo{
		ImageURL:  "127.0.0.1/img/foto.png",
		MAC:       "AA:BB:CC:DD:EE:FF",
		Timestamp: now,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	var respBody models.Photo
	err := json.Unmarshal(rec.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.NotZero(t, respBody.ID)
	assert.Equal(t, payload.MAC, respBody.MAC)
	assert.Equal(t, payload.ImageURL, respBody.ImageURL)
	assert.True(t, respBody.Timestamp.Equal(now))

	var saved models.Photo
	err = db.First(&saved, respBody.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, respBody.ID, saved.ID)
	assert.Equal(t, respBody.ImageURL, saved.ImageURL)
	assert.Equal(t, payload.MAC, saved.MAC)
}
