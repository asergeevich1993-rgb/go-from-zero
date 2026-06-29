package main

import "fmt"

func groupByLen(words []string) map[int][]string {
	result := map[int][]string{}

	for _, word := range words {
		length := len([]rune(word))
		result[length] = append(result[length], word)
	}

	return result
}

func main() {

	result := groupByLen([]string{"го", "это", "круто", "я"})
	fmt.Print(result)
}
