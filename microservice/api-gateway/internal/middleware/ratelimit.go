package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var requests = make(map[string]int)

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		requests[ip]++
		if requests[ip] > 20 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}