package entity

import (
	"github.com/google/uuid"
	"time"
)

type BaseEntity struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	DeviceID   string    `json:"device_id" gorm:"type:varchar(255);not null;index" validate:"required"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;not null;index" validate:"required"`
	ReceivedAt time.Time `json:"received_at" gorm:"type:timestamp;not null"`
}
