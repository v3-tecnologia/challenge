package database

import (
	"testing"

	"github.com/google/uuid"
	"github.com/igorlopes88/desafio-v3/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRegisterGps(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Gps{})
	gps, _ := entity.NewGps(uuid.New(), "00:1A:2B:3C:4D:5E", -23.853345114882476, -45.28845610372027, "2025-06-30 10:00:01")
	gpsDB := NewGps(db)

	err = gpsDB.Register(gps)
	assert.Nil(t, err)

	var dataFound entity.Gps
	err = db.First(&dataFound, "id = ?", gps.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, gps.ID, dataFound.ID)
	assert.Equal(t, gps.User, dataFound.User)
	assert.Equal(t, gps.MacAddress, dataFound.MacAddress)
	assert.Equal(t, gps.Latitude, dataFound.Latitude)
	assert.Equal(t, gps.Longitude, dataFound.Longitude)
	assert.Equal(t, gps.TimeStamp, dataFound.TimeStamp)
}
