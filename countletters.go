package main

import (
	"fmt"
	"strings"
)

func main() {

	text := ("арозаупаланалапуазора")

	letters := strings.Split(text, "")
	counter := map[string]int{}

	for _, letter := range letters {
		counter[letter] = counter[letter] + 1
	}
	for letter, count := range counter {
		fmt.Println(letter, ":", count)
	}
}
