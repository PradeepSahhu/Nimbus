package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yashpatil74/nimbus/internal/api/handlers"
	"github.com/yashpatil74/nimbus/internal/api/routes"
	"github.com/yashpatil74/nimbus/internal/db"
	"github.com/yashpatil74/nimbus/internal/repository"
	"github.com/yashpatil74/nimbus/internal/services"
)

func main() {

	// Setup
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

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
	UserRepository := repository.NewUserRepository(datab)

	// Services
	AuthService := services.NewAuthService(UserRepository)

	// Handlers
	genericHandler := handlers.NewGenericHandler()
	authHandler := handlers.NewAuthHandler(AuthService)

	router := gin.Default()
	apiRoute := router.Group("/api")
	{
		routes.SetupGenericRoutes(apiRoute, genericHandler)
		routes.SetupAuthRoutes(apiRoute, authHandler)
	}

	router.Run(os.Getenv("PORT"))

}
