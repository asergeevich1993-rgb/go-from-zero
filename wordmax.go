package main

import (
	"fmt"
	"strings"
)

func main() {

	products := map[string]int{}
	text := "арбуз банан арбуз груша банан арбуз"
	countproducts := strings.Split(text, " ")
	maxcount := 0
	for _, v := range countproducts {
		products[v] = products[v] + 1
	}
	for countproducts, count := range products {
		fmt.Println(countproducts, ":", count)

		if count > maxcount {
			maxcount = count
			fmt.Println("Самое частое слово ", countproducts, ":", maxcount)
		}
	}

}
