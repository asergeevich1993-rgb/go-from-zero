package main

import "fmt"

func main() {

	fruits := map[string]int{"яблоки": 10, "бананы": 5, "груши": 7}

	for fruit, count := range fruits {
		fmt.Println(fruit, ":", count)
	}

	delete(fruits, "бананы")

	for fruit, count := range fruits {
		fmt.Println("Вывод: ", fruit, ":", count)
	}
	delete(fruits, "Арбузы")

	for fruit, count := range fruits {
		fmt.Println("Вывод: ", fruit, ":", count)
	}
}
