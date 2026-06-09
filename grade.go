package main

import (
	"fmt"
)

func main() {
	var grade int

	fmt.Print("Введи оценку и получи описание: ")
	fmt.Scan(&grade)
	if grade == 5 {
		fmt.Println("Отлично!")
	} else if grade == 4 {
		fmt.Println("Хорошо!")
	} else if grade == 3 {
		fmt.Println("Удовлетворительно!")
	} else if grade == 2 {
		fmt.Println("неуд!")
	} else if grade == 1 {
		fmt.Println("отчислен!")
	} else {
		fmt.Println("нет такой оценки!")
	}
}
