package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "User Logged In successfully"})
}
