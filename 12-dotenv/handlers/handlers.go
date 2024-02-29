package handlers

import (
	Database "github.com/Celesca/12-dotenv/db"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// GET /products
func GetProducts(c *fiber.Ctx) error {

	var products []Database.Product
	Database.Db.Find(&products)
	return c.JSON(products)

}

// GET: /products/:id
func GetProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")

	var product Database.Product
	result := Database.Db.First(&product, productID) // productID is the number
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {

	var product Database.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	Database.Db.Create(&product)
	return c.JSON(product)
}
func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var product Database.Product
	result := Database.Db.First(&product, productID)
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

	Database.Db.Save(&product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var product Database.Product
	result := Database.Db.First(&product, productID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	Database.Db.Delete(&product)
	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user Database.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	user.Password = string(hashPassword)

	Database.Db.Create(&user)
	return c.JSON(user)
}
