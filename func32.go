package main

import "fmt"

func groupByLength(words []string) map[int][]string {
	counter := map[int][]string{}
	for _, word := range words {
		length := len([]rune(word))
		counter[length] = append(counter[length], word)
	}
	return counter

}

func main() {

	result := groupByLength([]string{"го", "это", "круто", "я", "ты"})
	fmt.Print(result)

}
