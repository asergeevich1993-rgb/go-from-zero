package main

import "fmt"

func main() {
	var sum float64
	fmt.Print("Введите сумму покупки: ")
	fmt.Scan(&sum)
	var fee float64

	if sum < 1000 {
		fee = 0
	} else if sum < 5000 {
		fee = 5
	} else if sum < 10000 {
		fee = 10
	} else {
		fee = 15
	}
	discontSum := sum * fee / 100
	total := sum - discontSum

	fmt.Printf("Сумма без скидки: %.2f руб.\n", sum)
	fmt.Printf("Скидка: %.0f%%\n", fee)
	fmt.Printf("Сумма скидки: %.2f руб.\n", discontSum)
	fmt.Printf("Итого к оплате: %.2f руб.\n", total)
}
