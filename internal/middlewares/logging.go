package middlewares

import (
	"log/slog"
	"time"

	"github.com/KaiRibeiro/challenge/internal/logs"
	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		logs.Logger.Info("request started",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("clientIP", c.ClientIP()),
		)

		c.Next()

		latency := time.Since(start).Milliseconds()

		logs.Logger.Info("request completed",
			slog.Int("status", c.Writer.Status()),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("clientIP", c.ClientIP()),
			slog.Int64("latency_ms", latency),
		)
	}
}
