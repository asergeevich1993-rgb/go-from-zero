package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "го го это круто го это"
	words := strings.Split(text, " ")

	counter := map[string]int{}

	for _, word := range words {
		counter[word]++
	}
	for key, w := range counter {
		fmt.Println(key, ":", w)
	}

}
