package api

import (
	"blockchain-project/blockchain" // Use the correct import path

	"github.com/gofiber/fiber/v2"
)

func setupQuoteRoutes(app *fiber.App) {
	app.Get("/api/swap/quote", func(c *fiber.Ctx) error {
		sellToken := c.Query("sellToken")
		buyToken := c.Query("buyToken")
		amount := c.Query("amount")               // This could be sellAmount or buyAmount, decide based on your API design
		isSellAmount := c.Query("type") == "sell" // Example logic to determine if it's sell or buy amount

		if sellToken == "" || buyToken == "" || amount == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required query parameters (sellToken, buyToken, amount)"})
		}

		quoteResponse, err := blockchain.GetSwapQuote(sellToken, buyToken, amount, isSellAmount)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(quoteResponse)
	})
}
