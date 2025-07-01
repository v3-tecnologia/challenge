package database

import (
	"github.com/igorlopes88/desafio-v3/internal/entity"
	"gorm.io/gorm"
)

type Photo struct {
	DB *gorm.DB
}

func NewPhoto(db *gorm.DB) *Photo {
	return &Photo{DB: db}
}

func (g *Photo) Register(photo *entity.Photo) error {
	return g.DB.Create(photo).Error
}