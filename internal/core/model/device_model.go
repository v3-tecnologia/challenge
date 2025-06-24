package model

import "time"

type Device struct {
	MacAddress string    `json:"mac_address"`
	Timestamp  time.Time `json:"timestamp"`
}
