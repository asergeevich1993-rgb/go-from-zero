package main

import (
	"fmt"
	"os"
)

func main() {

	data := []byte("Привет,файл!")
	os.WriteFile("hello.txt", data, 0644)

	content, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}

	fmt.Println(string(content))

}
