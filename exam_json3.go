package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Phone struct {
	Brand string
	Model string
	Price int
}

func main() {
	phone := []Phone{
		{Brand: "iPhone", Model: "13ProMax", Price: 400},
		{Brand: "Samsung", Model: "Galaxy", Price: 450},
		{Brand: "Nokia", Model: "Lumia", Price: 400}}

	data, err := json.Marshal(phone)
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}
	err = os.WriteFile("Phones.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи: ", err)
		return
	}
	fmt.Println("Saved")
	var phones []Phone = []Phone{}

	jsonData, err := os.ReadFile("Phones.json")
	if err != nil {
		fmt.Println("Ошибка чтения: ", err)
		return
	}
	err = json.Unmarshal(jsonData, &phones)
	if err != nil {
		fmt.Println("Ошибка десериализации: ", err)
		return
	}
	for _, p := range phones {
		fmt.Println(p.Brand, "", p.Model, ":", p.Price)
	}
}
