package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	file, err := os.Create("app.log")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	// set log output
	log.SetOutput(file)

	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("GET /")

		return c.SendString("GET /")
	})

	app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}

}
