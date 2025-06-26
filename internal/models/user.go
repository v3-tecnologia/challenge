package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Username  string    `gorm:"not null" json:"username" binding:"required"`
	Email     string    `gorm:"not null;unique" json:"email" binding:"required,email"`
	Password  string    `gorm:"not null" json:"password" binding:"required,min=8"`
	Roles     []Role    `gorm:"many2many:user_roles;" json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
