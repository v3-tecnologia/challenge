package dtos

type ErrorResponseDTO struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
