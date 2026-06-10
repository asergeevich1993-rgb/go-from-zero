package main

import "fmt"

func main() {

	result := 1
	for i := 1; i <= 10; i++ {
		result = result * 2
		fmt.Println("Вывод; ", result)
	}
}
