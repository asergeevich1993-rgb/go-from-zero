package main

import "fmt"

func printAny(v any) {
	fmt.Println(v)
}

func main() {
	printAny(42)
	printAny("Го го это круто")

	Fruit{Color: "Красный", Name: "Манго"}
	printAny(Fruit)
}
