package main

import "fmt"

func main() {
	var a int
	var b int
	var operation string

	for {
		fmt.Print("Введите первое число: ")
		fmt.Scan(&a)

		fmt.Print("Операция (+, -, *, / или 'выход'): ")
		fmt.Scan(&operation)

		if operation == "выход" {
			fmt.Println("Программа завершена.")
			break
		}

		fmt.Print("Введите второе число: ")
		fmt.Scan(&b)

		if operation == "+" {
			fmt.Println("Результат:", a+b)
		} else if operation == "-" {
			fmt.Println("Результат:", a-b)
		} else if operation == "*" {
			fmt.Println("Результат:", a*b)
		} else if operation == "/" {
			if b == 0 {
				fmt.Println("Ошибка: деление на ноль")
			} else {
				fmt.Println("Результат:", a/b)
			}
		} else {
			fmt.Println("Неизвестная операция")
		}
	}
}
