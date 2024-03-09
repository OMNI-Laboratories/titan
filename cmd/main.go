package main

import (
	"fmt"
	"log"

	"github.com/OMNI-Laboratories/titan/pkg/api"

	"github.com/OMNI-Laboratories/titan/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Retrieve the API key from environment variables
	apiKey := config.Config("OX_API_KEY")
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
