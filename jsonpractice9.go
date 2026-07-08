package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	cities := map[string][]string{"Москва": {"Кремль", "ВДНХ", "Парк Горького"},
		"Питер": {"Эрмитаж", "Петропавловка", "Разводные мосты"}}

	data, err := json.Marshal(cities)
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}
	err = os.WriteFile("cities.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи: ", err)
		return
	}
	fmt.Println("Saved")

	jdata, err := os.ReadFile("cities.json")
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}
	var c map[string][]string = map[string][]string{}
	err = json.Unmarshal(jdata, &c)
	if err != nil {
		fmt.Println("Ошибка десериализации: ", err)
		return
	}
	for key, value := range c {
		fmt.Println(key, ": ", value)
	}

}
