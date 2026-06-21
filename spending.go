package main

import "fmt"

func main() {

	wallet := map[string]int{"Артур": 6000, "Мария": 3000, "Иван": 7000}
	var name string
	var spend int

	for {
		fmt.Println("Кто? ")
		fmt.Scan(&name)
		if name == "стоп" {
			break
		}
		fmt.Print("Сколько: ")
		fmt.Scan(&spend)

		balance, exists := wallet[name]

		if exists && balance >= spend {
			wallet[name] = balance - spend
			fmt.Print(name, ":", wallet[name], "\n")
		} else if exists && balance < spend {
			fmt.Print("Недостаточно средств у ", name, ":", balance, "\n")
		} else {
			fmt.Print("такого человека нет\n")
		}

	}
	for name, spendbalance := range wallet {
		fmt.Println(name, ":", spendbalance)
	}

}
