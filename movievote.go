package main

import "fmt"

func main() {
	movievote := map[string]string{"Артур": "", "Мария": "", "Иван": ""}

	count := map[string]int{}
	maxVotes := 0
	winner := ""
	var movie string
	var name string
	for i := 0; i < 3; i++ {

		fmt.Print("Кто: ")
		fmt.Scan(&name)
		fmt.Print("За что? ")
		fmt.Scan(&movie)

		namevote, exists := movievote[name]

		if exists && namevote == "" {
			movievote[name] = movie
			fmt.Println(name, ":", movievote[name])
			count[movie] = count[movie] + 1
		} else if exists {
			fmt.Println("Уже голосовал")
		} else {
			fmt.Println("Такого участника нет")
		}

	}

	for movie, votes := range count {
		if votes > maxVotes {
			maxVotes = votes
			winner = movie
		}
	}
	fmt.Println("Результаты: ")
	for name, vote := range movievote {
		fmt.Println(name, ":", vote)

	}
	fmt.Println("Победил: ", winner, maxVotes, " голосов")

}
