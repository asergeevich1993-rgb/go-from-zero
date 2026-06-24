package main

import "fmt"

func isPalindrome(nums []int) bool {

	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}

	return true

}

func main() {

	result1 := isPalindrome([]int{1, 2, 3, 2, 1})
	fmt.Print(result1)

	result2 := isPalindrome([]int{1, 2, 3, 4, 5})
	fmt.Print(result2)
}
