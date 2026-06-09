package main

import "fmt"

func main() {
	name := "Artur"
	age := 33
	city := "Moscow"

	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Город:", city)

	fmt.Printf("Меня зовут %s, мне %d года, я из города %s.\n", name, age, city)
}
