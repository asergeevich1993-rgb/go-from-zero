package main

import "fmt"

func filer(nums []int) []int {

	evennums := []int{}
	for _, value := range nums {
		if value%2 != 0 {
			continue
		}
		evennums = append(evennums, value)
	}
	return evennums

}

func main() {

	result := filer([]int{1, 2, 3, 4, 5, 6})
	fmt.Println("фильтр четных: ", result)
}
