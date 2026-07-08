package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	product := map[string]int{"карандаши": 100, "ручки": 200, "ластики": 250}

	data, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println(string(data))
	err = os.WriteFile("Products.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи: ", err)
		return
	}

	fmt.Println("Saved")

	jdata, err1 := os.ReadFile("Products.json")
	if err1 != nil {
		fmt.Println("Ошибка чтения: ", err)
		return
	}

	p := map[string]int{}

	err1 = json.Unmarshal(jdata, &p)
	if err1 != nil {
		fmt.Println("Ошибка десериализации: ", err1)
		return
	}
	fmt.Println(p)

}
