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
	data, err := os.ReadFile("Cars.json")
	if err != nil {
		fmt.Println("файл не найден: ", err)
		return
	}
	var car []Car

	err = json.Unmarshal(data, &car)
	if err != nil {
		fmt.Println("Ошибка десериализации: ", err)
		return
	}
	for _, c := range car {
		fmt.Println(c.Brand, "", c.Model, "", c.Year)
	}
}
