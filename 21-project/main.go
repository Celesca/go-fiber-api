package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageURL string `json:"imageURL"`
}

func generateFakeData() []Product {
	products := make([]Product, 0)
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 10; i++ {
		product := Product{
			ID:       i,
			Name:     fmt.Sprintf("Product %d", i),
			Price:    rand.Intn(100) + 1,
			ImageURL: fmt.Sprintf("https://plus.unsplash.com/premium_photo-1663954865317-3e2c288cf5be?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"),
		}

		products = append(products, product)
	}

	return products
}

func main() {

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Response().Header.Set("Access-Control-Allow-Origin", "*")
		c.Response().Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Response().Header.Set("Access-Control-Allow-Header", "Origin, Content-Type, Accept")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}

		return c.Next()

	})

	app.Get("/api/products", func(c *fiber.Ctx) error {
		products := generateFakeData()
		return c.JSON(products)
	})

	log.Fatal(app.Listen(":8000"))

}
