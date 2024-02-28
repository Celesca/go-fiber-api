package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getUsers(c *fiber.Ctx) error {
	name := c.Query("name")

	users := []User{
		{
			ID: "B01", Name: "Aef",
		},
		{ID: "B02", Name: "Bef"},
	}

	if name != "" {
		filteredUsers := []User{}
		for _, user := range users {
			if user.Name == name {
				filteredUsers = append(filteredUsers, user)
			}
		}

		return c.JSON(filteredUsers)
	}

	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := User{
		ID:   id,
		Name: "Aef",
	}

	return c.JSON(user)
}

func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	response := map[string]string{
		"message": "User created successfully",
	}
	return c.JSON(response)
}

func main() {
	app := fiber.New()

	// Query parameter (c.QueryParams("paramName")) ?name=

	app.Get("/users", getUsers)

	// Specific Path parameter (c.Params("name"))

	app.Get("/users/:id", getUser)

	// Post request body in JSON c.Body() or c.FormValue("paramName")

	app.Post("/users/", createUser)

	// c.BodyParser(&data) -> from JSON to go struct

	app.Listen(":3000")

}
