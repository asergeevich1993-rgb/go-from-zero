package main

import "fmt"

func main() {

	nums := []int{5, 7, 8, 10, 12, 105}
	count := 0
	/*evensum := 0
	oddsum := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			evensum = evensum + nums[i]
		} else {
			oddsum = oddsum + nums[i]
		}
	}
	fmt.Println("Четная сумма: ", evensum)
	fmt.Println("нечетная сумма: ", oddsum)
	*/

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			count++
		} else {
		}

	}
	fmt.Println("количество четных: ", count)
}
