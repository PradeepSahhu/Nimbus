package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/api/handlers"
	"github.com/yashpatil74/nimbus/internal/api/routes"
	"github.com/yashpatil74/nimbus/internal/db"
)

func main() {

	// Database
	datab, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer datab.Close()

	if err := db.RunMigrations(datab); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

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
