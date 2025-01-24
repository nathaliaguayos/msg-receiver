package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nathaliaguayos/msg-receiver/internal/services"
	"net/http"
)

// JWTHandler is the interface that provides JWT handling methods.
//
//counterfeiter:generate . JWTHandler
type JWTHandler interface {
	GenerateToken(c *gin.Context)
}

type jwtHandler struct {
	jwtService services.JWTService
}

// NewJWTHandler creates a new JWTHandler.
func NewJWTHandler(jwtService services.JWTService) JWTHandler {
	return &jwtHandler{
		jwtService: jwtService,
	}
}

// GenerateToken generates a JWT token.
// Params: c *gin.Context - the request context
func (h *jwtHandler) GenerateToken(c *gin.Context) {
	var request struct {
		UserID string `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.jwtService.GenerateToken(request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
