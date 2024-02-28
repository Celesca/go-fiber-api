package main

import (
	"encoding/json"
	"fmt"
)

type MenuItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func main() {

	jsonString := `{
		"name": "Americano",
		"description": "Espresso with hot water",
		"price": 10.5
	}`

	var menuItem MenuItem
	err := json.Unmarshal([]byte(jsonString), &menuItem)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Printf("%s (%s) : %.2f\n", menuItem.Name, menuItem.Description, menuItem.Price)
}
