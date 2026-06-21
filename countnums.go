package main

import "fmt"

func main() {

	nums := []int{5, 3, 5, 5, 7, 3}

	counter := map[int]int{}

	for _, value := range nums {
		counter[value] = counter[value] + 1
	}

	for key, count := range counter {
		fmt.Println(key, ":", count)
	}

}
