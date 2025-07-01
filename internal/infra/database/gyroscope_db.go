package database

import (
	"github.com/igorlopes88/desafio-v3/internal/entity"
	"gorm.io/gorm"
)

type Gyroscope struct {
	DB *gorm.DB
}

func NewGyroscope(db *gorm.DB) *Gyroscope {
	return &Gyroscope{DB: db}
}

func (g *Gyroscope) Register(gyroscope *entity.Gyroscope) error {
	return g.DB.Create(gyroscope).Error
}