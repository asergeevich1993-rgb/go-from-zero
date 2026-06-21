package main

import "fmt"

func main() {

	stock := map[string]int{"ручки": 100, "тетради": 50}
	var name string
	var count int
	fmt.Print("Что пришло: ")
	fmt.Scan(&name)
	fmt.Print("Сколько: ")
	fmt.Scan(&count)

	stock[name] = stock[name] + count
	fmt.Println(name, ":", stock[name])
	for name, v := range stock {
		fmt.Println(name, ":", v)
	}
}
