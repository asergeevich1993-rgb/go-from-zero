package main

import "fmt"

func main() {

	atm := map[string]int{"Артур": 10000, "Мария": 5000, "Иван": 15000}
	var name string
	var sum int
	var action int

	for i := 0; i <= 5; i++ {

		fmt.Print(" Кто? ")
		fmt.Scan(&name)
		fmt.Print("Снять (1) Пополнить (2): ")
		fmt.Scan(&action)
		if action == 1 {
			fmt.Print("Сумма: ")
			fmt.Scan(&sum)
			balance, exists := atm[name]
			if exists && atm[name] >= sum {
				atm[name] = balance - sum
				fmt.Print("Баланс ", name, ":", atm[name])

			} else if exists && atm[name] < sum {
				fmt.Println("Баланса недостаточно: ", atm[name])

			} else {
				fmt.Println("Клиент не найден")
			}
		} else if action == 2 {
			fmt.Print("Сумма: ")
			fmt.Scan(&sum)
			balance, exists := atm[name]

			if exists {
				atm[name] = balance + sum
				fmt.Println("Баланс: ", atm[name])
			} else {
				fmt.Println("Клиент не найден")
			}

		}
	}
	for name, balance2 := range atm {
		fmt.Println("Итог; ", name, ":", balance2)
	}
}
