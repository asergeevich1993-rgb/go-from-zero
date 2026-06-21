package main

import "fmt"

func main() {

	stock := map[string]int{"ручки": 100, "тетради": 50}

	var ware string
	var count int

	for i := 0; i < 5; i++ {
		fmt.Print("Товар: ")
		fmt.Scan(&ware)
		fmt.Print("Колличество ")
		fmt.Scan(&count)

		remainder, exists := stock[ware]
		if exists && remainder+count < 0 {
			fmt.Println("товара недостаточно")
		} else if exists && remainder+count >= 0 {
			stock[ware] = stock[ware] + count
			fmt.Print(ware, ":", stock[ware], "\n")
		} else {
			stock[ware] = stock[ware] + count
			fmt.Print(ware, ":", stock[ware], "\n")
		}

	}
	for name, value := range stock {
		fmt.Println(name, ":", value)

	}

}
