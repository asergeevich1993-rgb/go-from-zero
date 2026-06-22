package main

import "fmt"

func main() {

	votes := []string{"дюна", "барби", "дюна", "оппенгеймер", "барби", "дюна"}
	var winner string
	maxVotes := 0
	film := map[string]int{}

	for _, vote := range votes {
		film[vote]++
	}

	for name, value := range film {
		fmt.Println(name, ":", value)
		if value > maxVotes {
			maxVotes = value
			winner = name
		}
	}
	fmt.Println("победитель: ", winner, "(", maxVotes, " голосов)")

}
