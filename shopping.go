package main

import "fmt"

func main() {
	/*list := []string{"Хлеб", "Молоко", "Яйца"}

	fmt.Println("Список покупок:")
	for i := 0; i < len(list); i++ {
		fmt.Println(i+1, list[i])
	}*/
	num := []int{10, 20, 30, 40, 50}
	sum := 0
	for i := 0; i < len(num); i++ {
		sum = sum + num[i]
	}
	fmt.Print("Сумма: ", sum)

}
