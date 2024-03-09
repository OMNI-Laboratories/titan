package api

import (
	"titan/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the routes for the API group
func SetupRoutes(app fiber.Router) {
	// Group API routes
	api := app.Group("/")
	// Define a sub-group, if needed, for versioning or further categorization
	// v1 := api.Group("/v1")
	api.Get("/home", handlers.Home)
}
