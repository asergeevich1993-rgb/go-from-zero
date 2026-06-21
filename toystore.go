package main

import "fmt"

func main() {

	toystore := map[string]int{"мишки": 10, "куклы": 5, "машинки": 8}
	var toy string
	var count int

	for i := 1; i <= 5; i++ {
		fmt.Print("Что продаем? ")
		fmt.Scan(&toy)
		fmt.Print("Сколько? ")
		fmt.Scan(&count)

		stock, exists := toystore[toy]

		if exists && toystore[toy] >= count {
			toystore[toy] = stock - count
			fmt.Print("продано: ", count, " Осталось: ", toystore[toy])

		} else if exists && toystore[toy] < count {
			fmt.Print("Недостаточно! На складе: ", stock)

		} else {
			fmt.Print("нет в наличии!")
		}
		fmt.Println(" ")
	}

}
