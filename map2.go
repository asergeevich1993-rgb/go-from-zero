package main

import "fmt"

func main() {

	scores := map[string]int{}

	scores["алебра"] = 5
	scores["русский"] = 4

	for name, v := range scores {
		fmt.Println(name, ":", v)
	}
}
