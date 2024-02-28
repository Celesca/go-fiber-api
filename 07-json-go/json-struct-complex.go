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

	menu := []MenuItem{
		{
			Name:        "Americano",
			Description: "Espresso with hot water",
			Price:       10.0,
			Ingredients: []Ingredients{
				{
					Name:     "Espresso",
					Quantity: 1,
					Unit:     "shot",
				},
				{
					Name:     "Hot water",
					Quantity: 1,
					Unit:     "cup",
				},
			},
		},
		{
			Name:        "Latte",
			Description: "Espresso with steamed milk",
			Price:       20.0,
			Ingredients: []Ingredients{
				{
					Name:     "Espresso",
					Quantity: 1,
					Unit:     "shot",
				},
				{
					Name:     "Steamed milk",
					Quantity: 1,
					Unit:     "cup",
				},
			},
		},
	}

	bytes, _ := json.Marshal(menu)
	fmt.Println(bytes)
}
