package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Car struct {
	Brand string
	Model string
	Year  int
}

func main() {

	data, err := os.ReadFile("Car.json")
	if err != nil {
		fmt.Println("Ошибка чтения: ", err)
		return
	}
	var c Car
	err = json.Unmarshal(data, &c)
	if err != nil {
		fmt.Println("Ошибка форматирования из JSON в Go: ", err)
		return
	}
	fmt.Println(c.Brand, " ", c.Model, " ", c.Year)

}
