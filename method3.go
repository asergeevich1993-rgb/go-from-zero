package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

func (i Book) info() {
	fmt.Println("Книга: ", i.Title, "страниц:", i.Pages)
}

func (i Book) IsThick() bool {
	return i.Pages > 300
}

func main() {
	book1 := Book{Title: "Война и мир", Pages: 1225}
	book1.info()
	fmt.Println("больше трехсот страниц: ", book1.IsThick())
}
