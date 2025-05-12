package domain

import (
	"errors"
	"regexp"
)

var (
	ErrInValidMACGyroscope = errors.New("invalid MAC address")
)

type Device struct {
	ID string
}

func NewDevice(id string) (*Device, error) {
	if !IsValidMAC(id) {
		return nil, ErrInValidMACGyroscope
	}
	return &Device{ID: id}, nil
}

func IsValidMAC(id string) bool {
	return regexp.MustCompile(`^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$`).MatchString(id)
}
