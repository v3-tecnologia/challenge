package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
	rate     rate.Limit
	burst    int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		limiters: make(map[string]*rate.Limiter),
		rate:     r,
		burst:    b,
	}
}

func (i *IPRateLimiter) getLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.limiters[ip]
	if !exists {
		limiter = rate.NewLimiter(i.rate, i.burst)
		i.limiters[ip] = limiter
	}

	return limiter
}

func (i *IPRateLimiter) Allow(ip string) bool {
	return i.getLimiter(ip).Allow()
}

// RateLimitMiddleware cria um middleware de rate limiting
// rate: número de requisições por segundo permitidas
// burst: número máximo de requisições em rajada
func RateLimitMiddleware(requestsPerSecond int, burst int) gin.HandlerFunc {
	limiter := NewIPRateLimiter(rate.Limit(requestsPerSecond), burst)

	// Goroutine para limpar limiters antigos (cleanup a cada 10 minutos)
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			limiter.mu.Lock()
			// Remove limiters que não foram usados recentemente
			for ip, l := range limiter.limiters {
				// Se o limiter permite todas as requisições, significa que não foi usado recentemente
				if l.TokensAt(time.Now()) == float64(limiter.burst) {
					delete(limiter.limiters, ip)
				}
			}
			limiter.mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !limiter.Allow(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "muitas requisições. tente novamente em alguns segundos",
				"code":  "RATE_LIMIT_EXCEEDED",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// TelemetryRateLimitMiddleware - Rate limiting específico para endpoints de telemetria
// Permite 100 requisições por minuto (mais restritivo para endpoints críticos)
func TelemetryRateLimitMiddleware() gin.HandlerFunc {
	// Em modo de teste, use rate limiting mais permissivo
	if gin.Mode() == gin.TestMode {
		return RateLimitMiddleware(1000, 100) // Muito permissivo para testes
	}
	return RateLimitMiddleware(100, 10) // 100 req/min com burst de 10 para produção
}

// AuthRateLimitMiddleware - Rate limiting para endpoints de autenticação
// Permite 10 tentativas por minuto (mais restritivo para prevenir ataques)
func AuthRateLimitMiddleware() gin.HandlerFunc {
	// Em modo de teste, use rate limiting mais permissivo
	if gin.Mode() == gin.TestMode {
		return RateLimitMiddleware(1000, 100) // Muito permissivo para testes
	}
	return RateLimitMiddleware(10, 5) // 10 req/min com burst de 5 para produção
}

// GlobalRateLimitMiddleware - Rate limiting global mais permissivo
func GlobalRateLimitMiddleware() gin.HandlerFunc {
	return RateLimitMiddleware(200, 20) // 200 req/min com burst de 20
}
