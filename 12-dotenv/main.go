package main

import (
	"log"

	Db "github.com/Celesca/12-dotenv/db"
	HandlerController "github.com/Celesca/12-dotenv/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	Db.InitDatabase()
	app := fiber.New()

	app.Get("/products", HandlerController.GetProducts)
	app.Get("/products/:id", HandlerController.GetProductByID)
	app.Post("/products", HandlerController.CreateProduct)
	app.Put("/products/:id", HandlerController.UpdateProduct)
	app.Delete("/products/:id", HandlerController.DeleteProduct)

	// Users
	app.Post("/users", HandlerController.CreateUser)

	log.Fatal(app.Listen(":3000"))
}
