package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

// RateLimiter creates a rate limiter for each client IP
func RateLimiter(limit rate.Limit) gin.HandlerFunc {
	// Map to store rate limiters for each client IP
	limiterMap := make(map[string]*rate.Limiter)
	var mu sync.Mutex

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		// Ensure thread-safe access to the map
		mu.Lock()
		if _, exists := limiterMap[clientIP]; !exists {
			limiterMap[clientIP] = rate.NewLimiter(limit, int(limit))
		}
		mu.Unlock()

		// Check if the client is allowed to proceed
		if !limiterMap[clientIP].Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			c.Abort()
			return
		}

		// Continue with the request
		c.Next()
	}
}
