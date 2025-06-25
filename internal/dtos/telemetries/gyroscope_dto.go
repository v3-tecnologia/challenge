package dtos

type CreateGyroscopeDto struct {
	X *float64 `json:"x" binding:"required"`
	Y *float64 `json:"y" binding:"required"`
	Z *float64 `json:"z" binding:"required"`
}
