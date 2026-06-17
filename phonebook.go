package main

import "fmt"

func main() {
	book := map[string]string{
		"Артур": "+79991234567",
		"Мария": "+79261234567",
	}

	var name string
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	phone, exists := book[name]
	if exists {
		fmt.Println("Телефон:", phone)
	} else {
		fmt.Println("Не найден")
	}
}
