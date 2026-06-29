package main

import (
	"fmt"
)

func groupbyletterA(words []string) map[int][]string {
	result := map[int][]string{}

	for _, word := range words {
		count := 0
		letters := []rune(word)
		for _, letter := range letters {
			if string(letter) == "а" /*strings.Contains("а", string(letter))*/ {
				count++
			}
		}
		result[count] = append(result[count], word)
	}
	return result
}

func main() {
	result := groupbyletterA([]string{"арбуз", "банан", "апельсин", "киви"})
	fmt.Print(result)
}
