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
	c := []Car{{Brand: "Toyota", Model: "Camry", Year: 2020},
		{Brand: "BMW", Model: "X5", Year: 2021}}

	jsonData, err := json.Marshal(c)
	if err != nil {
		fmt.Print("Ошибка сереализации: ", err)
		return
	}

	err = os.WriteFile("Cars.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка записи: ", err)
		return
	}
	fmt.Println("saved")

}
