package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserID   uint   `json:"userid"`
	Username string `json:"username"`
}

var secretKey = []byte("mysecretkey")

func loginHandler(c *fiber.Ctx) error {

	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})

	}

	if user.Username != "admin" || user.Password != "password" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Credentials"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = 1
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not login"})
	}

	return c.JSON(fiber.Map{"token": tokenString})

}

func protectedHandler(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenString := ""

	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Author"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Token"})
	}

	// Validated Token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Token"})
	}

	userID := uint(claims["userId"].(float64))
	username := claims["username"].(string)

	return c.JSON(fiber.Map{"message": fmt.Sprintf("Proctected Route accesses by user ID %d, username: %s", userID, username)})
}

func main() {

	app := fiber.New()

	app.Post("/login", loginHandler)

	app.Get("/admin", protectedHandler)

	log.Fatal(app.Listen(":3000"))

}
