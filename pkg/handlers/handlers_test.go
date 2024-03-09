package handlers

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	// Setup Fiber app
	app := fiber.New()

	// Define the route identical to how it's defined in your actual app
	app.Get("/home", Home)

	// Create a new HTTP request for the "/home" route
	req := httptest.NewRequest("GET", "/home", nil)

	// Perform the request
	resp, _ := app.Test(req)

	// Assert the status code is what you expect: 200 OK
	assert.Equal(t, 200, resp.StatusCode)

	// Read the response body
	body, _ := ioutil.ReadAll(resp.Body)

	// Convert the body to a string and assert its value
	assert.Equal(t, "Welcome to Titan!!!", string(body))
}
