package main

import "fmt"

func main() {

	library := map[string]bool{"Войнаимир": true, "Капитанскаядочка": true, "ЕвгенийОнегин": true}

	var book string

	for i := 0; i < 4; i++ {
		fmt.Println("Что берем? ")
		fmt.Scan(&book)

		bookbool, exists := library[book]
		if exists && bookbool == true {
			fmt.Println("Выдана!")
			library[book] = false
		} else if exists && bookbool == false {
			fmt.Println("нет в наличии")

		} else {
			fmt.Println("нет такой книги!")
		}
	}

	for book, bookbalance := range library {
		fmt.Println(" остаток ", book, ":", bookbalance)
	}
}
