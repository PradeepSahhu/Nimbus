package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/api/handlers"
)

func SetupGenericRoutes(router *gin.RouterGroup, handler *handlers.GenericHandler) {
	router.GET("/ping", handler.Ping)
}
