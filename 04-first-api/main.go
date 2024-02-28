package main

import (
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Post request")
	})

	app.Listen(":3000")

}
