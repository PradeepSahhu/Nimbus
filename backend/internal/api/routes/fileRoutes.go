package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/api/handlers"
)

func SetupFileRoutes(router *gin.RouterGroup, fileHandler *handlers.FileHandler) {
	routerGroup := router.Group("/file")
	{
		routerGroup.POST("/upload", nil)
	}
}
