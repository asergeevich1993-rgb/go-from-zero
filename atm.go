package main

import "fmt"

func main() {

	balance := 10000
	var operation1 int
	var sum int

	for i := 1; i <= 5; i++ {
		fmt.Println("операция", i, "Снять (1) Пополнить (2)")
		fmt.Scan(&operation1)

		if operation1 == 1 {
			fmt.Println("Уточните сумму: ", "Баланс: ", balance)
			fmt.Scan(&sum)
			if sum > balance {
				fmt.Println("недостаточно средсвт")
			} else {
				balance -= sum
				fmt.Println("Баланс: ", balance)
			}

		} else if operation1 == 2 {
			fmt.Println("Уточните сумму: ", sum, "Баланс: ", balance)
			fmt.Scan(&sum)
			balance += sum
			fmt.Println("Баланс: ", balance)
		} else {
			fmt.Println("неизвестная команда")

		}

	}

	fmt.Println("итоговый баланс: ", balance)

}
