package dtos

import "time"

type BaseDTO struct {
	MacAddress string `json:"mac_address" binding:"required"`
	Timestamp  int64  `json:"timestamp" binding:"required"`
}

func (b *BaseDTO) GetTimestamp() time.Time {
	return time.Unix(b.Timestamp, 0)
}
