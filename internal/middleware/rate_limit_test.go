package middleware

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestRateLimiter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Define the rate limit (e.g., 2 requests per second)
	limit := rate.Limit(2)

	// Create a new Gin router with the rate limiter middleware
	router := gin.New()
	router.Use(RateLimiter(limit))

	// Define a test route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("X-Real-IP", "127.0.0.1") // Mock client IP

	// Create a ResponseRecorder to capture the response
	w := httptest.NewRecorder()

	// First request should pass
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Second request should pass
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Third request should be blocked due to rate limit
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusTooManyRequests, w.Code)
}

func TestRateLimiter_ConcurrentRequests(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Define the rate limit (2 requests per second)
	limit := rate.Limit(2)

	// Create a new Gin router with the rate limiter middleware
	router := gin.New()
	router.Use(RateLimiter(limit))

	// Define a test route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Number of concurrent requests
	numRequests := 5
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Shared response codes storage
	responseCodes := make([]int, numRequests)

	// Simulate concurrent requests
	for i := 0; i < numRequests; i++ {
		go func(index int) {
			defer wg.Done()
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			req.Header.Set("X-Real-IP", "127.0.0.1") // Same mock client IP

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			responseCodes[index] = w.Code
		}(i)
	}

	wg.Wait()

	// Count how many requests were allowed vs. rate limited
	var successCount, rateLimitedCount int
	for _, code := range responseCodes {
		if code == http.StatusOK {
			successCount++
		} else if code == http.StatusTooManyRequests {
			rateLimitedCount++
		}
	}

	// At most 2 requests should succeed due to the rate limit
	assert.LessOrEqual(t, successCount, 2)
	// The remaining requests should be rate-limited
	assert.GreaterOrEqual(t, rateLimitedCount, 3)
}

func TestRateLimiter_ResetAfterDuration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Define the rate limit (1 request per second)
	limit := rate.Limit(1)

	// Create a new Gin router with the rate limiter middleware
	router := gin.New()
	router.Use(RateLimiter(limit))

	// Define a test route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("X-Real-IP", "127.0.0.1") // Mock client IP

	// First request should pass
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Second request should be blocked
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusTooManyRequests, w.Code)

	// Wait for the rate limiter to reset
	time.Sleep(1 * time.Second)

	// Third request should pass after waiting
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
