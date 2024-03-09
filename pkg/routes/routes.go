package routes

import (
	"github.com/OMNI-Laboratories/titan/pkg/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/")

	api.Get("/home", handlers.Home)
	api.Get("/swap/quote", handlers.SwapQuoteHandler)
}
