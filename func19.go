package main

import "fmt"

func merge(a, b []int) []int {
	ab := []int{}

	for _, value := range a {
		ab = append(ab, value)
	}
	for _, value2 := range b {
		ab = append(ab, value2)
	}

	return ab

}

func main() {

	result := merge([]int{1, 2, 3}, []int{4, 5})

	fmt.Print(result)

}
