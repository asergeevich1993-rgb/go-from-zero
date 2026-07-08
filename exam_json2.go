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

	var phone Phone

	jsonData, err := os.ReadFile("phone.json")
	if err != nil {
		fmt.Println("Ошибка чтения: ", err)
		return
	}
	err = json.Unmarshal(jsonData, &phone)
	if err != nil {
		fmt.Println("Ошибка десериализации: ", err)
		return
	}
	fmt.Println(phone.Brand, ":", phone.Model, "", phone.Price)

}
