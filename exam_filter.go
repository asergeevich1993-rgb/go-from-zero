package main

import "fmt"

func main() {

	nums := []int{7, 12, 5, 8, 3, 10}
	var evennums []int
	count := 0
	for _, v := range nums {
		if v%2 == 0 {
			count++
			evennums = append(evennums, v)
		}
	}
	fmt.Println("исходный", nums)
	fmt.Println("Четные: ", evennums)
	fmt.Print("Колличество: ", count)

}
