package main

import (
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Get request")
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Post request")
	})

	app.Put("/", func(c *fiber.Ctx) error {
		return c.SendString("Put request")
	})

	app.Delete("/", func(c *fiber.Ctx) error {
		return c.SendString("Delete request")
	})

	app.Listen(":3000")

}
