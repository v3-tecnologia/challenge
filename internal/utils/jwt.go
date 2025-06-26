package utils

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	jwtSecret []byte
	once      sync.Once
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Roles  []string  `json:"roles"`
	jwt.RegisteredClaims
}

func getJWTSecret() []byte {
	once.Do(func() {
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "secret-key-256-bits"
		}
		jwtSecret = []byte(secret)
	})
	return jwtSecret
}

func GenerateAccessToken(userID uuid.UUID, email string, roles []string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

func GenerateRefreshToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(30 * 24 * time.Hour).Unix() // 30 dias
	claims["iat"] = time.Now().Unix()
	claims["jti"] = uuid.New().String()

	return token.SignedString(getJWTSecret())
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ValidateRefreshToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
