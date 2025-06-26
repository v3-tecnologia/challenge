package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"telemetry-api/internal/dtos/requests"
	"telemetry-api/internal/services"
	"telemetry-api/internal/utils"
)

type AuthHandler struct {
	service *services.AuthService
	logger  *zap.Logger
}

func NewAuthHandler(service *services.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{service: service, logger: logger}
}

func (h *AuthHandler) CreateUser(c *gin.Context) {
	var user requests.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("failed to bind JSON", zap.Error(err))
		errs := utils.TranslateValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	response, err := h.service.CreateUser(&user)
	if err != nil {
		h.logger.Error("failed to create user", zap.Error(err))

		// Verificar se é um erro específico da aplicação
		if utils.IsAppError(err, utils.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": utils.TranslateAppError(err)})
			return
		}

		if utils.IsAppError(err, utils.ErrRoleNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": utils.TranslateAppError(err)})
			return
		}

		// Erro genérico
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao criar usuário"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest requests.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		h.logger.Error("failed to bind JSON", zap.Error(err))
		errs := utils.TranslateValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	response, err := h.service.Login(&loginRequest)
	if err != nil {
		h.logger.Error("failed to login", zap.Error(err))

		// Verificar se é um erro específico da aplicação
		if utils.IsAppError(err, utils.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": utils.TranslateAppError(err)})
			return
		}

		// Erro genérico
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao fazer login"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var refreshRequest requests.RefreshTokenRequest

	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		h.logger.Error("failed to bind JSON", zap.Error(err))
		errs := utils.TranslateValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	response, err := h.service.RefreshToken(&refreshRequest)
	if err != nil {
		h.logger.Error("failed to refresh token", zap.Error(err))

		// Verificar se é um erro específico da aplicação
		if utils.IsAppError(err, utils.ErrRefreshTokenExpired) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": utils.TranslateAppError(err)})
			return
		}

		// Erro genérico
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao renovar token"})
		return
	}

	c.JSON(http.StatusOK, response)
}
