package main

import "fmt"

func countDups(words []string) map[string]int {

	counter := map[string]int{}
	for _, value := range words {
		counter[value]++
	}
	/*for key, value := range counter {
		counter[key] = value
	}*/
	return counter

}

func main() {

	result := countDups([]string{"го", "это", "го", "круто", "это", "го"})

	fmt.Print(result)

}
