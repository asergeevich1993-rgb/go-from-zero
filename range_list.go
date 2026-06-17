package main

import "fmt"

func main() {

	/*cities := []string{"москва", "питер", "казань", "сочи"}

	for i, v := range cities {
		fmt.Println(i+1, v)
	}*/
	nums := []int{4, 8, 15, 16, 23, 42}
	/*sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	fmt.Print(sum)*/
	fmt.Print(nums)
	for i := range nums {
		nums[i] = 0
	}
	fmt.Print(nums)
}
