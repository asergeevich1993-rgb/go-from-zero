package main

import "fmt"

func filterEven(nums []int) []int {
	Evennums := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			Evennums = append(Evennums, v)
		}

	}
	return Evennums
}

func main() {
	result := filterEven([]int{3, 8, 1, 12, 4})

	fmt.Println("Четные: ", result)
}
