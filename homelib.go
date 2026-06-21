package main

import "fmt"

func main() {
	var book string

	homelib := map[string]bool{"1984": true, "МастериМаргарита": true, "Тритоварища": false}

	for {
		fmt.Print("Что ищем?: ")
		fmt.Scan(&book)
		if book == "выход" {
			break
		}
		bookbool, exists := homelib[book]

		if exists && bookbool == true {
			fmt.Println("держи книгу!")
			homelib[book] = false

		} else if exists {
			fmt.Println("Отдал другу!")
		} else {
			fmt.Println("нет такой книги")
		}
	}

	for book, bookbalance := range homelib {
		if bookbalance {
			fmt.Print(book, ": Дома\n")
		} else {
			fmt.Print(book, ": у друга\n")
		}
	}
}
