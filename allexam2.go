package main

import "fmt"

func countFruits(fruits []string) map[string]int {
	counter := map[string]int{}

	for _, fruit := range fruits {
		counter[fruit]++
	}
	for name, frts := range counter {
		counter[name] = frts
	}
	return counter
}

func main() {

	result := countFruits([]string{"яблоко", "банан", "яблоко", "груша", "банан", "яблоко"})
	fmt.Print(result)

}
