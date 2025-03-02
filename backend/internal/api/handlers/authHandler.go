package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/services"
)

type AuthHandler struct {
	AuthService   *services.AuthService
	FolderService *services.FolderService
}

func NewAuthHandler(authService *services.AuthService, folderService *services.FolderService) *AuthHandler {
	return &AuthHandler{
		AuthService:   authService,
		FolderService: folderService,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.AuthService.Register(requestBody.Username, requestBody.Email, requestBody.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.FolderService.CreateUserFolder(user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User Registered successfully"})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.AuthService.Login(requestBody.Email, requestBody.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
