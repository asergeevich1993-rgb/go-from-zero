package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	var x int
	sum := 0
	for i := 0; i < 5; i++ {
		fmt.Print("Введите индекс: ")
		fmt.Scan(&x)
		nums = append(nums, x)
		sum = sum + x + nums[i]
	}
	fmt.Print(nums, "\n")
	fmt.Print("Сумма: ", sum)
}
