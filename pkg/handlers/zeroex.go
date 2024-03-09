package handlers

import (
	"github.com/OMNI-Laboratories/titan/pkg/blockchain"
	"github.com/gofiber/fiber/v2"
)

// SwapQuoteHandler handles requests for swap quotes
func SwapQuoteHandler(c *fiber.Ctx) error {
	sellToken := c.Query("sellToken")
	buyToken := c.Query("buyToken")
	amount := c.Query("amount")
	isSellAmount := c.Query("type") == "sell"

	if sellToken == "" || buyToken == "" || amount == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required query parameters (sellToken, buyToken, amount)"})
	}

	quoteResponse, err := blockchain.GetSwapQuote(sellToken, buyToken, amount, isSellAmount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(quoteResponse)
}
