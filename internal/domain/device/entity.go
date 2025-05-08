package device

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
	if !isValidMAC(id) {
		return nil, ErrInValidMACGyroscope
	}
	return &Device{ID: id}, nil
}

func isValidMAC(id string) bool {
	return regexp.MustCompile(`^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$`).MatchString(id)
}
