package main

import "fmt"

func main() {

	warehouse := map[string]int{"яблоки": 100, "бананы": 50, "груши": 75}

	var warename string
	var ware int
	fmt.Print("Что отгружаем? ")
	fmt.Scan(&warename)
	fmt.Print("Сколько?")
	fmt.Scan(&ware)

	stock, exists := warehouse[warename]
	if exists && warehouse[warename] >= ware {
		warehouse[warename] = stock - ware
		fmt.Print("Осталось ", warename, ":", warehouse[warename])
	} else if exists && warehouse[warename] < ware {
		fmt.Print("недостаточно в наличии", stock, "шт")
	} else {
		fmt.Print("Товар не найден")
	}

}
