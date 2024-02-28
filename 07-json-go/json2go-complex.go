package main

import (
	"encoding/json"
	"fmt"
)

type MenuItem struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Price       float64       `json:"price"`
	Ingredients []Ingredients `json:"ingredients"`
}

type Ingredients struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

func main() {

	jsonString := `[
	{
		"name": "Americano",
		"description": "Espresso with hot water",
		"price": 10.5,
		"ingredients": [
			{
				"name": "Espresso",
				"quantity": 1,
				"unit": "shot"
			},
			{
				"name": "Hot water",
				"quantity": 1,
				"unit": "cup"
			}
		]
	},
	{
		"name": "Late",
		"description": "Late with hot water",
		"price": 10.5,
		"ingredients": [
			{
				"name": "Espresso",
				"quantity": 1,
				"unit": "shot"
			},
			{
				"name": "Hot water",
				"quantity": 1,
				"unit": "cup"
			}
		]
	}
]`

	var menuItem []MenuItem
	err := json.Unmarshal([]byte(jsonString), &menuItem)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	for _, menumenuItem := range menuItem {
		fmt.Printf("%s (%s) : %.2f\n", menumenuItem.Name, menumenuItem.Description, menumenuItem.Price)
	}

}
