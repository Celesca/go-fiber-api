package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson: "_id, omitempty"`
	Name     string             `bson: "name"`
	Email    string             `bson: "email"`
	Password string             `bson: "password"`
}

func main() {
	app := fiber.New()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		collection := client.Database("demogofiber").Collection("users")

		password := "password123"
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			log.Fatal(err)
		}

		user := User{
			Name:     "Sirasit",
			Email:    "SiraCore",
			Password: string(hashPassword),
		}

		result, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted document with ID :", result.InsertedID)

		return c.SendString("Created User successfully.")
	})

	app.Listen(":3000")

}
