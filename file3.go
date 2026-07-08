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
	p := Person{Name: "Артур", Age: 33}

	jData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	fmt.Println(string(jData))
	err = os.WriteFile("person.json", jData, 0644)
	if err != nil {
		fmt.Println("Ошибка записи: ", err)
	}
	fmt.Println("Saved!")
}
