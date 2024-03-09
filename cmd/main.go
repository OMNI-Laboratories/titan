package main

import (
	"log"

	"github.com/OMNI-Laboratories/titan/pkg/config"
	"github.com/OMNI-Laboratories/titan/pkg/routes"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Retrieve the API key from environment variables
	oxAPIKey := config.Config("OX_API_KEY")
	if oxAPIKey == "" {
		log.Fatal("OX_API_KEY must be set in .env")
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Initialize your API routes within the group
	routes.SetupRoutes(app)

	log.Println("Listening to port 3000")
	app.Listen(":3000")
}
