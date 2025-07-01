package dto

import "github.com/igorlopes88/desafio-v3/pkg/entity"

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterGyroscopeInput struct {
	UserId     entity.ID `json:"user"`
	MacAddress string    `json:"mac_address"`
	XAxis      float64   `json:"x"`
	YAxis      float64   `json:"y"`
	ZAxis      float64   `json:"z"`
	TimeStamp  string    `json:"timestamp"`
}

type RegisterGpsInput struct {
	UserId     entity.ID `json:"user"`
	MacAddress string    `json:"mac_address"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	TimeStamp  string    `json:"timestamp"`
}

type RegisterPhotoInput struct {
	UserId     entity.ID `json:"user"`
	MacAddress string    `json:"mac_address"`
	Image      string    `json:"image"`
	TimeStamp  string    `json:"timestamp"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}