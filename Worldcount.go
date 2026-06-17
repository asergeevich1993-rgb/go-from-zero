package main

import "fmt"

func main() {

	grades := []int{5, 4, 3, 5, 5, 4, 3, 5, 4}

	counter := map[int]int{}

	for i := 0; i < len(grades); i++ {
		grade := grades[i]
		counter[grade] = counter[grade] + 1

	}
	for grade, count := range counter {
		fmt.Printf("оценка %d: %d раз\n", grade, count)
	}
}
