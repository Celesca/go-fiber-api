package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model          // has ID and createdAt, updatedAt
	Name        string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Description string
}

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique; not null"`
	Password  string `gorm:"not null"`
}

// GET /products
func GetProducts(c *fiber.Ctx) error {

	var products []Product
	db.Find(&products)
	return c.JSON(products)

}

// GET: /products/:id
func GetProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")

	var product Product
	result := db.First(&product, productID) // productID is the number
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {

	var product Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	db.Create(&product)
	return c.JSON(product)
}
func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var product Product
	result := db.First(&product, productID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	// Parse the body to product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	db.Save(&product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var product Product
	result := db.First(&product, productID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	db.Delete(&product)
	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

var db *gorm.DB

func initDatabase() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPost := os.Getenv("DB_PORT")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPost + ")/" + dbName + "?parseTime=true"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	db.AutoMigrate(&Product{}, &User{})
}

func main() {

	initDatabase()
	app := fiber.New()

	app.Get("/products", GetProducts)
	app.Get("/products/:id", GetProductByID)
	app.Post("/products", CreateProduct)
	app.Put("/products/:id", UpdateProduct)
	app.Delete("/products/:id", DeleteProduct)

	log.Fatal(app.Listen(":3000"))
}
