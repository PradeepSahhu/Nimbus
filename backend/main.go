package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/controllers"
	"github.com/yashpatil74/nimbus/internal/routes"
	// _ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

func main() {

	router := gin.Default()

	// Controllers
	routeController := controllers.NewRouterController()

	apiRoute := router.Group("/api")
	{
		routes.SetupRoutes(apiRoute, routeController)
	}

	router.Run(":8080")
}
