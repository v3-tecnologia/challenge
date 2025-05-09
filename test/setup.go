package test

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
	"v3-backend-challenge/src/api"
	"v3-backend-challenge/src/model"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "sqlite",
		DSN:        "file::memory:?cache=shared",
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	err = db.AutoMigrate(&model.GPS{}, model.Gyroscope{}, model.Photo{})
	if err != nil {
		panic("failed to migrate test database")
	}
	return db
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	telemetryApi := api.TelemetryApi{}
	telemetryApi.RegisterRoutes(r, db)
	return r
}
