package main

import "fmt"

func sumValues(m map[string]int) int {
	sum := 0
	for _, value := range m {
		sum = sum + value
	}

	return sum
}

func main() {

	result := sumValues(map[string]int{"ручки": 100, "тетради": 50, "карандаши": 200})

	fmt.Print("сумма: ", result)
}
