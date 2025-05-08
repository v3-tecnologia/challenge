package device

import (
	"errors"
	"regexp"
)

type Device struct {
	ID string
}

func NewDevice(id string) (*Device, error) {
	if !isValidMAC(id) {
		return nil, errors.New("invalid MAC address")
	}
	return &Device{ID: id}, nil
}

func isValidMAC(id string) bool {
	return regexp.MustCompile(`^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$`).MatchString(id)
}
