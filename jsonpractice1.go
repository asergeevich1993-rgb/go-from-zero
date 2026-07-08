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
	c := Car{Brand: "Toyota", Model: "Camry", Year: 2020}

	jsonData, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Ошибка записи в JSON: ", err)
		return
	}
	fmt.Println(string(jsonData))
	err = os.WriteFile("Car.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка создания файла: ", err)
		return
	}

	fmt.Println("Save complited")

}
