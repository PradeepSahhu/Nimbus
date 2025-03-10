package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/api/handlers"
)

func SetupAuthRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	routerGroup := router.Group("/auth")
	{
		routerGroup.POST("/register", authHandler.Register)
		routerGroup.POST("/login", authHandler.Login)
	}
}
