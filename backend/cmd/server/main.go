package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/api/handlers"
	"github.com/yashpatil74/nimbus/internal/api/routes"
)

func main() {
	// Repositories

	// Services

	// Handlers
	genericHandler := handlers.NewGenericHandler()
	authHandler := handlers.NewAuthHandler()

	router := gin.Default()
	apiRoute := router.Group("/api")
	{
		routes.SetupGenericRoutes(apiRoute, genericHandler)
		routes.SetupAuthRoutes(apiRoute, authHandler)
	}

	router.Run(":8080")

}
