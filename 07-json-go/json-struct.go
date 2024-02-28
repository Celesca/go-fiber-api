package main

import (
	"encoding/json"
	"fmt"
)

type Coffee struct {
	Menu     string
	Price    int
	Quantity int
}

func main() {

	myCoffee := Coffee{"Cappucion", 20, 2}

	bytes, _ := json.Marshal(myCoffee)

	fmt.Println(string(bytes))

}
