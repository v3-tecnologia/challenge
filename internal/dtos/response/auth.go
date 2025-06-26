package response

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	User             UserResponse `json:"user"`
	AccessToken      string       `json:"access_token"`
	RefreshToken     string       `json:"refresh_token"`
	ExpiresIn        int64        `json:"expires_in"`
	RefreshExpiresIn int64        `json:"refresh_expires_in"`
	TokenType        string       `json:"token_type"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
