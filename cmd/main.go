package main

import (
	"fmt"
	"log"
	"os"
	"titan/pkg/api" // Adjust the import path according to your project structure

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the API key from environment variables
	oxAPIKey := os.Getenv("OX_API_KEY")
	if oxAPIKey == "" {
		log.Fatal("OX_API_KEY must be set in .env")
	}

	app := fiber.New()
	app.Use(cors.New())

	// Initialize your API routes within the group
	api.SetupRoutes(app)

	fmt.Println("Listening to port 3000")
	app.Listen(":3000")
}
