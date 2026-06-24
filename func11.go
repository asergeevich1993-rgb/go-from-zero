package main

import "fmt"

func main() {

	resulteven := groupEvenOdd([]int{1, 2, 3, 4, 5, 6, 7})

	fmt.Println("Четные/нечетные: ", resulteven)

}

func groupEvenOdd(nums []int) map[string]int {
	evenodd := map[string]int{}
	for _, value := range nums {
		if value%2 == 0 {
			evenodd["even"]++
		} else {
			evenodd["odd"]++
		}
	}

	return evenodd

}
