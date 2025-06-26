package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"telemetry-api/internal/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "cabeçalho de autorização é obrigatório"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "formato do cabeçalho de autorização inválido"})
			c.Abort()
			return
		}

		token := bearerToken[1]
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("roles", claims.Roles)
		c.Next()
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "roles não encontrados no contexto"})
			c.Abort()
			return
		}

		userRoles, ok := roles.([]string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "formato de roles inválido"})
			c.Abort()
			return
		}

		for _, userRole := range userRoles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "permissões insuficientes"})
		c.Abort()
	}
}

func RequireAnyRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoles, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "roles não encontrados no contexto"})
			c.Abort()
			return
		}

		userRolesList, ok := userRoles.([]string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "formato de roles inválido"})
			c.Abort()
			return
		}

		for _, requiredRole := range roles {
			for _, userRole := range userRolesList {
				if userRole == requiredRole {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "permissões insuficientes"})
		c.Abort()
	}
}
