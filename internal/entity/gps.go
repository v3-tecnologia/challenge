package entity

import (
	"time"

	"github.com/igorlopes88/desafio-v3/pkg/entity"
)

type Gps struct {
	ID         entity.ID `json:"id" gorm:"primaryKey"`
	User       entity.ID `json:"user" gorm:"not null"`
	MacAddress string    `json:"mac_address" gorm:"not null"`
	Latitude   float64   `json:"latitude" gorm:"not null"`
	Longitude  float64   `json:"longitude" gorm:"not null"`
	TimeStamp  time.Time `json:"timestamp" gorm:"not null"`
}

func NewGps(userId entity.ID, mac_address string, latitude, longitude float64, time string) (*Gps, error) {
	newtime, err := SrtToTime(time)
	if err != nil {
		return nil, err
	}
	gps := &Gps{
		ID:         entity.NewID(),
		User:       userId,
		MacAddress: mac_address,
		Latitude:   latitude,
		Longitude:  longitude,
		TimeStamp:  newtime,
	}
	err = gps.Validate()
	if err != nil {
		return nil, err
	}
	return gps, nil
}

func (g *Gps) Validate() error {
	if g.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(g.ID.String()); err != nil {
		return ErrInvalidID
	}
	if g.User.String() == "" {
		return ErrInvalidUser
	}
	if _, err := entity.ParseID(g.User.String()); err != nil {
		return ErrInvalidUser
	}
	if g.MacAddress == "" {
		return ErrMacIsRequired
	}
	return nil
}
