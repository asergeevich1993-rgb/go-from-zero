package main

import "fmt"

func main() {
	secret := 67
	var guess int
	won := false

	fmt.Println("Угадай число: ")
	for i := 1; i <= 7; i++ {
		fmt.Println("Номер попытки: ", i, " попыток из 7")
		fmt.Scan(&guess)
		if guess == secret {
			fmt.Println("Угадал!")
			won = true
			break
		} else if guess > secret {
			fmt.Println("Меньше!")
		} else {
			fmt.Println("Больше!")
		}

	}
	if !won {
		fmt.Println("Проиграл число было: ", secret)
	}

}
