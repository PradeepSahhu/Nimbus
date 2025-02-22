package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/controllers"
)

func SetupRoutes(router *gin.RouterGroup, routeController *controllers.RouterController) {
	router.GET("/health", routeController.CheckHealth)
}
