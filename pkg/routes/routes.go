package routes

import (
	"log"

	"github.com/OMNI-Laboratories/titan/pkg/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	log.Println("Setting up routes")
	api := app.Group("/")

	api.Get("/home", handlers.Home)
	api.Get("/swap/quote", handlers.SwapQuoteHandler)
}
