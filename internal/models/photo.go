package models

import "time"

type Photo struct {
	ID         uint64    `gorm:"primaryKey" json:"id,omitempty"`
	MAC        string    `gorm:"index;not null" json:"mac" validate:"required"`
	ImageURL   string    `gorm:"" json:"image_url"`
	Recognized bool      `gorm:"default:false" json:"recognized"`
	Timestamp  time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
