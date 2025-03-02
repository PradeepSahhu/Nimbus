package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
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
	FolderRepository := repository.NewFolderRepository(datab)
	UserRepository := repository.NewUserRepository(datab)
	FileRepository := repository.NewFileRepository(datab)

	// Services
	FolderService := services.NewFolderService(FolderRepository, `../../data/files`)
	AuthService := services.NewAuthService(UserRepository)
	FileService := services.NewFileService(FileRepository, FolderRepository, `../../data/files`)

	// Handlers
	genericHandler := handlers.NewGenericHandler()
	authHandler := handlers.NewAuthHandler(AuthService, FolderService)
	fileHandler := handlers.NewFileHandler(FileService)

	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"0.0.0.0"} // Your frontend URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	apiRoute := router.Group("/api")
	{
		routes.SetupGenericRoutes(apiRoute, genericHandler)
		routes.SetupAuthRoutes(apiRoute, authHandler)
		routes.SetupFileRoutes(apiRoute, fileHandler)
	}

	router.Run(os.Getenv("PORT"))
}
