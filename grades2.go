package main

import "fmt"

func main() {

	grades := []int{5, 4, 5, 5, 4, 3, 5, 4}
	count := map[int]int{}
	for _, v := range grades {
		count[v]++
	}
	for key, countmap := range count {
		fmt.Println(key, ":", countmap)
	}

}
