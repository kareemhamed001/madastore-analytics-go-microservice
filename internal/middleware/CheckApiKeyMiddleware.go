package middleware

import (
	"madastore/analytics/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckApiKeyMiddleware(c *gin.Context) {
	apiKey := c.GetHeader("X-API-KEY")
	expectedApiKey := config.Load().ApiKey // Replace with your actual expected API key

	if apiKey != expectedApiKey {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Next()
}
