package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		duration := time.Since(start)
		requestID, _ := c.Get(RequestIDKey)
		requestIDStr, _ := requestID.(string)
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		log.Info().
			Str("request_id", requestIDStr).
			Str("method", c.Request.Method).
			Str("path", path).
			Int("status", c.Writer.Status()).
			Dur("duration", duration).
			Str("client_ip", c.ClientIP()).
			Msg("http_request")
	}
}
