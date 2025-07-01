package database

import (
	"github.com/igorlopes88/desafio-v3/internal/entity"
	"gorm.io/gorm"
)

type Gps struct {
	DB *gorm.DB
}

func NewGps(db *gorm.DB) *Gps {
	return &Gps{DB: db}
}

func (g *Gps) Register(gps *entity.Gps) error {
	return g.DB.Create(gps).Error
}