package entity

import (
	"time"

	"github.com/igorlopes88/desafio-v3/pkg/entity"
)

type Gyroscope struct {
	ID         entity.ID `json:"id" gorm:"primaryKey"`
	User       entity.ID `json:"user" gorm:"not null"`
	MacAddress string    `json:"mac_address" gorm:"not null"`
	XAxis      float64   `json:"x" gorm:"not null"`
	YAxis      float64   `json:"y" gorm:"not null"`
	ZAxis      float64   `json:"z" gorm:"not null"`
	TimeStamp  time.Time `json:"timestamp" gorm:"not null"`
}

func NewGyroscope(userId entity.ID, mac_address string, x, y, z float64, time string) (*Gyroscope, error) {
	newtime, err := SrtToTime(time)
	if err != nil {
		return nil, err
	}
	gyro := &Gyroscope{
		ID:         entity.NewID(),
		User:       userId,
		MacAddress: mac_address,
		XAxis:      x,
		YAxis:      y,
		ZAxis:      z,
		TimeStamp:  newtime,
	}
	err = gyro.Validate()
	if err != nil {
		return nil, err
	}
	return gyro, nil
}

func (g *Gyroscope) Validate() error {
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
