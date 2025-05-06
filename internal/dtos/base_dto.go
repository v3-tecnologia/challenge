package dtos

type BaseDTO struct {
	MacAddress string `json:"mac_address" binding:"required"`
}
