package main

import "fmt"

func main() {

	nums := []int{12, -5, 33, 0, 18, -2}
	min := nums[0]
	max := nums[0]

	for _, v := range nums {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	fmt.Println("минимум: ", min)
	fmt.Println("максимум: ", max)

}
