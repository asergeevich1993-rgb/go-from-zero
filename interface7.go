package main

import "fmt"

func printAny(v any) {
	fmt.Println(v)

}

type Book struct {
	Title string
}

func main() {

	t := Book{Title: "Го"}
	printAny(t)
}
