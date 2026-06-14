package main

import (
	"log"
	"os"

	"go-matrix-api/internal/infrastructure/client"
	"go-matrix-api/internal/infrastructure/handler"
	"go-matrix-api/internal/middleware"
	"go-matrix-api/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	
	_ "go-matrix-api/docs"
)

// @title Interseguro Matrix API
// @version 1.0
// @description API for matrix rotation and QR factorization
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your.email@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Escribe "Bearer " seguido de tu token JWT.
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load, falling back to environment variables")
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Dependency Injection
	nodeClient := client.NewNodeClient()
	matrixUsecase := usecase.NewMatrixUsecase(nodeClient)
	
	matrixHandler := handler.NewMatrixHandler(matrixUsecase)
	authHandler := handler.NewAuthHandler()

	// Setup routes
	api := app.Group("/api")
	
	// Auth route
	api.Get("/auth/token", authHandler.GenerateToken)

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Protected routes
	api.Post("/matrix", middleware.JWTProtected(), matrixHandler.ProcessMatrix)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}