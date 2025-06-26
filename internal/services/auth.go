package services

import (
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"telemetry-api/internal/dtos/requests"
	"telemetry-api/internal/dtos/response"
	"telemetry-api/internal/models"
	"telemetry-api/internal/utils"
)

type AuthService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewAuthService(db *gorm.DB, logger *zap.Logger) *AuthService {
	return &AuthService{db: db, logger: logger}
}

func (s *AuthService) CreateUser(user *requests.CreateUserRequest) (response.UserResponse, error) {
	var existingUser models.User
	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		s.logger.Warn("attempt to create user with existing email", zap.String("email", user.Email))
		return response.UserResponse{}, utils.ErrEmailAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("failed to hash password", zap.Error(err))
		return response.UserResponse{}, err
	}

	// Get admin role
	var adminRole models.Role
	if err := s.db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		s.logger.Error("admin role not found", zap.Error(err))
		return response.UserResponse{}, utils.ErrRoleNotFound
	}

	userModel := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
		Roles:    []models.Role{adminRole},
	}

	if err := s.db.Create(&userModel).Error; err != nil {
		s.logger.Error("failed to create user in service", zap.Error(err))
		return response.UserResponse{}, err
	}

	// Extract role names for response
	var roles []string
	for _, role := range userModel.Roles {
		roles = append(roles, role.Name)
	}

	createUserResponse := response.UserResponse{
		ID:        userModel.ID,
		Username:  userModel.Username,
		Email:     userModel.Email,
		Roles:     roles,
		CreatedAt: userModel.CreatedAt,
	}

	return createUserResponse, nil
}

func (s *AuthService) Login(loginRequest *requests.LoginRequest) (response.LoginResponse, error) {
	var user models.User
	if err := s.db.Preload("Roles").Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		s.logger.Warn("login attempt with non-existent email", zap.String("email", loginRequest.Email))
		return response.LoginResponse{}, utils.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		s.logger.Warn("login attempt with invalid password", zap.String("email", loginRequest.Email))
		return response.LoginResponse{}, utils.ErrInvalidCredentials
	}

	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, roles)
	if err != nil {
		s.logger.Error("failed to generate access token", zap.Error(err))
		return response.LoginResponse{}, err
	}

	refreshTokenString, err := utils.GenerateRefreshToken()
	if err != nil {
		s.logger.Error("failed to generate refresh token", zap.Error(err))
		return response.LoginResponse{}, err
	}

	refreshToken := models.RefreshToken{
		Token:     refreshTokenString,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}

	if err := s.db.Create(&refreshToken).Error; err != nil {
		s.logger.Error("failed to save refresh token", zap.Error(err))
		return response.LoginResponse{}, err
	}

	return response.LoginResponse{
		User: response.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Roles:     roles,
			CreatedAt: user.CreatedAt,
		},
		AccessToken:      accessToken,
		RefreshToken:     refreshTokenString,
		ExpiresIn:        3600,    // 1 hora
		RefreshExpiresIn: 2592000, // 30 dias
		TokenType:        "Bearer",
	}, nil
}

func (s *AuthService) RefreshToken(refreshRequest *requests.RefreshTokenRequest) (response.RefreshTokenResponse, error) {
	isValid, err := utils.ValidateRefreshToken(refreshRequest.RefreshToken)
	if err != nil || !isValid {
		s.logger.Warn("invalid refresh token format", zap.Error(err))
		return response.RefreshTokenResponse{}, utils.ErrRefreshTokenExpired
	}

	// Check if refresh token exists in database and is not expired
	var refreshToken models.RefreshToken
	if err := s.db.Preload("User.Roles").Where("token = ? AND expires_at > ?", refreshRequest.RefreshToken, time.Now()).First(&refreshToken).Error; err != nil {
		s.logger.Warn("refresh token not found or expired", zap.Error(err))
		return response.RefreshTokenResponse{}, utils.ErrRefreshTokenExpired
	}

	// Extract role names
	var roles []string
	for _, role := range refreshToken.User.Roles {
		roles = append(roles, role.Name)
	}

	// Generate new access token
	accessToken, err := utils.GenerateAccessToken(refreshToken.User.ID, refreshToken.User.Email, roles)
	if err != nil {
		s.logger.Error("failed to generate access token", zap.Error(err))
		return response.RefreshTokenResponse{}, err
	}

	return response.RefreshTokenResponse{
		AccessToken: accessToken,
		ExpiresIn:   3600, // 1 hora
		TokenType:   "Bearer",
	}, nil
}
