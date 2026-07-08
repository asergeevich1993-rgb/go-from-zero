package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string

	Age int
}

func main() {
	p := Person{Name: "Артур", Age: 33}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	fmt.Println(string(jsonData))

}
