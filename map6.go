package main

import "fmt"

func main() {

	stock := map[string]int{"ручки": 100, "тетради": 100, "линейки": 40}
	var name string
	var count int

	fmt.Print("Что берете?: ")
	fmt.Scan(&name)
	fmt.Println("Сколько: ")
	fmt.Scan(&count)

	countproducts, exists := stock[name]
	if exists && countproducts >= count {
		stock[name] = countproducts - count
	} else if exists && countproducts < count {
		fmt.Print("Недостаточно\n")
	} else {
		fmt.Print("такого товара нет\n")
	}

	fmt.Println("Остаток: ", name, " ", stock[name])
}
