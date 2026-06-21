package main

import "fmt"

func main() {

	votes := []string{"Дюна", "Гладиатор", "Дюна", "Оппенгеймер", "Гладиатор", "Дюна"}

	counter := map[string]int{}

	var winner string
	countMax := 0

	for _, film := range votes {
		counter[film]++
	}

	for film, vote := range counter {
		fmt.Println(film, ":", vote)
		if vote > countMax {
			countMax = vote
			winner = film
		}
	}
	fmt.Print("победил: ", winner, "(", countMax, " раз(а) )")

}
