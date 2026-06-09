package main

import "fmt"

func main() {
	product := "молоко"
	price := 89.90
	quantity := 3
	taxRate := 0.20
	fmt.Println("Товар: ", product)
	fmt.Println("Цена: ", price)
	fmt.Println("Количество:", quantity)
	total := price * float64(quantity)
	Tax := total * taxRate
	All := total + Tax

	//fmt.Println("Сумма без НДС: ", total)
	//fmt.Println("Сумма НДС: ", Tax)
	//fmt.Println("Итого: ", All)

	fmt.Printf("Cумма без НДС: %.2f rub.\n", total)
	fmt.Printf("Сумма НДС: %.2f rub.\n", Tax)
	fmt.Printf("ИТОГО: %.2f rub.\n", All)

}
