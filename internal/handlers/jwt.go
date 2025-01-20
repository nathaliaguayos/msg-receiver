package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nathaliaguayos/msg-receiver/internal/services"
	"net/http"
)

type JWTHandler struct {
	jwtService services.JWTService
}

func NewJWTHandler(jwtService services.JWTService) *JWTHandler {
	return &JWTHandler{
		jwtService: jwtService,
	}
}

func (h *JWTHandler) GenerateToken(c *gin.Context) {
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
