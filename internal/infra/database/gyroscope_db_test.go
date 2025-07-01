package database

import (
	"testing"

	"github.com/google/uuid"
	"github.com/igorlopes88/desafio-v3/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRegisterGyroscope(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Gyroscope{})
	gyro, _ := entity.NewGyroscope(uuid.New(), "00:1A:2B:3C:4D:5E", 1, 1.1, -1.1, "2025-06-30 10:00:01")
	gyroDB := NewGyroscope(db)

	err = gyroDB.Register(gyro)
	assert.Nil(t, err)

	var dataFound entity.Gyroscope
	err = db.First(&dataFound, "id = ?", gyro.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, gyro.ID, dataFound.ID)
	assert.Equal(t, gyro.User, dataFound.User)
	assert.Equal(t, gyro.MacAddress, dataFound.MacAddress)
	assert.Equal(t, gyro.XAxis, dataFound.XAxis)
	assert.Equal(t, gyro.YAxis, dataFound.YAxis)
	assert.Equal(t, gyro.ZAxis, dataFound.ZAxis)
	assert.Equal(t, gyro.TimeStamp, dataFound.TimeStamp)
}
