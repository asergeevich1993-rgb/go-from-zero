package main

import "fmt"

func main() {

	nums := []int{10, 20, 30, 40, 50}
	var n int
	found := false
	fmt.Print("Число: ")
	fmt.Scan(&n)

	for i := 0; i < len(nums); i++ {
		if n == nums[i] {
			found = true
		} else {

		}
	}
	if found {
		fmt.Print("Угадал")
	} else {
		fmt.Print("неугадал")
	}
}
