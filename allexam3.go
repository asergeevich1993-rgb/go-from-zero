package main

import "fmt"

func groupbylen(words []string) map[int][]string {
	result := map[int][]string{}

	for _, word := range words {
		length := len([]rune(word))
		result[length] = append(result[length], word)
	}
	return result
}
func main() {

	result := groupbylen([]string{"го", "го", "это", "круто", "я"})
	fmt.Print(result)
}
