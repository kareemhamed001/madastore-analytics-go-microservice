package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckApiKeyMiddleware(expectedApiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey != expectedApiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}
