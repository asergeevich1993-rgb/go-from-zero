package main

import "fmt"

func main() {

	nums := []int{5, 7, 8, 10, 12, 105}
	count := 0

	for i := range nums {
		if nums[i]%2 == 0 {
			count++
		}

	}
	fmt.Print("Четных: ", count)

}
