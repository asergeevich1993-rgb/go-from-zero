package main

import "fmt"

func groupByFistLetter(words []string) map[string][]string {
	result := map[string][]string{}

	for _, word := range words {
		fistletter := string([]rune(word)[0])
		result[fistletter] = append(result[fistletter], word)
	}

	return result

}

func main() {

	result := groupByFistLetter([]string{"арбуз", "апельсин", "банан", "груша", "гранат"})

	fmt.Print(result)
}
