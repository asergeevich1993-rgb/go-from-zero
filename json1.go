package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	data, err := os.ReadFile("person.json")
	if err != nil {
		fmt.Print("Фаил не найден: ", err)
	}

	var p Person
	json.Unmarshal(data, &p)
	fmt.Println(p.Name, p.Age)
}
