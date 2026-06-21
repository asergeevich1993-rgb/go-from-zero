package main

import (
	"fmt"
	"strings"
)

func main() {

	counter := map[string]int{}
	text := "артур учит го артур пишет код го это круто"
	words := strings.Split(text, " ")

	for _, w := range words {
		counter[w]++
	}

	for words, count := range counter {
		fmt.Println(words, ":", count)
	}
}
