package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func helloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func TestHelloHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/", helloHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)

	// Check Status code

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check response body

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(body))
}

func TestMain(m *testing.M) {
	exitCode := m.Run()

	os.Exit(exitCode)
}
