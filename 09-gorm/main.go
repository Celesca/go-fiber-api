package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	ID          uint    `gorm: "primary_key"`
	Name        string  `gorm: "not null"`
	Price       float64 `gorm: "not null"`
	Description string
}

func main() {
	dsn := "root:12345@tcp(127.0.0.1:3306)/demo_ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/products", func(c *fiber.Ctx) error {
		var products []Product
		db.Find(&products)
		return c.JSON(products)
	})

	log.Fatal(app.Listen(":3000"))
}
