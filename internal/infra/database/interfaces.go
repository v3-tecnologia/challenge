package database

import "github.com/igorlopes88/desafio-v3/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type GyroscopeInterface interface{
	Register(gyroscope *entity.Gyroscope) error
}

type GpsInterface interface{
	Register(gps *entity.Gps) error
}

type PhotoInterface interface{
	Register(photo *entity.Photo) error
}