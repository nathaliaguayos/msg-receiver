package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nathaliaguayos/msg-receiver/internal/handlers"
	"github.com/nathaliaguayos/msg-receiver/internal/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

// Client represents a REST client.
type Client struct {
	Logger     *zerolog.Logger
	Router     *gin.Engine
	jwtHandler handlers.JWTHandler
}

// NewRestClient creates a new REST client.
func NewRestClient(log *zerolog.Logger, jwtHandler handlers.JWTHandler, rateLimit float64) (*Client, error) {
	if log == nil {
		return nil, errors.New("logger should not be null")
	}

	if jwtHandler == nil {
		return nil, errors.New("jwtHandler should not be null")
	}
	var instance = Client{
		Logger:     log,
		jwtHandler: jwtHandler,
	}

	router := gin.Default()
	router.Use(middleware.RateLimiter(rate.Limit(rateLimit)))
	log.Info().Int("rate_limit", int(rateLimit)).Msg("configured rate limit")
	router.POST("/token", jwtHandler.GenerateToken)

	instance.Router = router
	return &instance, nil
}
