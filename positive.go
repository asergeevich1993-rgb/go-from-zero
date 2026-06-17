package main

import "fmt"

func main() {

	nums := []int{-3, 5, -1, 8, 0, -2, 4}
	var posnums []int
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			count++
			posnums = append(posnums, nums[i])
		}

	}
	fmt.Print("Положительные: ", posnums, "\n")
	fmt.Print("Количество: ", count)
}
