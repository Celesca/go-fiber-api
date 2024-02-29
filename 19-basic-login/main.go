package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {

	app := fiber.New()

	credentials := map[string]string{
		"username1": "password1",
		"username2": "password2",
	}

	// Middleware

	auth := basicauth.New(basicauth.Config{
		Users: credentials,
	})

	app.Get("/protected", auth, func(c *fiber.Ctx) error {
		return c.SendString("Welcome to my pages.")
	})

	app.Listen(":3000")

}
