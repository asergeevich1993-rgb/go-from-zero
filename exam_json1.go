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
	phone := Phone{Brand: "Iphone", Model: "13ProMax", Price: 400}

	data, err := json.Marshal(phone)
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}

	err = os.WriteFile("Phone.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return

	}
	fmt.Println("Saved")

}
