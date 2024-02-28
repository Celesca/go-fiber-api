package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	Coffee := map[string]int{
		"Americano": 10,
		"Latte":     20,
		"Espresso":  30,
	}

	bytes, _ := json.Marshal(Coffee)
	fmt.Printf("Type of bytes %T\n", string(bytes))
	fmt.Println(string(bytes))
}
