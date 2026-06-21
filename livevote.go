package main

import "fmt"

func main() {

	allowed := map[string]bool{"Артур": false, "Мария": false, "Иван": false}
	counter := map[string]int{}
	var name string
	var film string
	var winner string
	maXcount := 0

	for {
		fmt.Print("Имя: ")
		fmt.Scan(&name)
		if name == "стоп" {
			break
		}
		fmt.Print("Фильм: ")
		fmt.Scan(&film)

		vote, exists := allowed[name]
		if exists && vote == false {
			allowed[name] = true
			fmt.Print(name, " проголосовал за ", film, "\n")
			counter[film]++
		} else if exists && vote == true {
			fmt.Print(name, " уже проголосовал\n")
		} else {
			fmt.Print("нет в списке голосующих\n")
		}
	}
	for film, votes := range counter {
		fmt.Println(film, ":", votes)
		if votes > maXcount {
			maXcount = votes
			winner = film
		}
	}
	fmt.Println(winner, " набрал ", maXcount, " голосов")

}
