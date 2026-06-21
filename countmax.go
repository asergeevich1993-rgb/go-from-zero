package main

import "fmt"

func main() {

	nums := []int{5, 3, 5, 3, 5, 7, 3}

	counter := map[int]int{}
	countMax := 0
	winner := 0
	for _, number := range nums {
		counter[number]++
	}

	for key, count := range counter {
		fmt.Println(key, ":", count)
		if countMax < count {
			countMax = count
			winner = key
		}
	}
	fmt.Println("Самое частое: ", winner, "(", countMax, " раз)")

}
