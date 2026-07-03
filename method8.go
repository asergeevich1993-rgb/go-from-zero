package main

import (
	"fmt"
)

type Book struct {
	Title     string
	Author    string
	Available bool
}

func (p Book) Print() {
	fmt.Println("книга: ", p.Title, "Автор: ", p.Author, "Наличие: ", p.Available)

}
func (b *Book) Borrow() bool {
	if b.Available == true {
		b.Available = false
		return true
	}
	return false
}
func (r *Book) Return() bool {
	if r.Available == false {
		r.Available = true
		return true
	}
	return false
}

func main() {
	library := []Book{
		{Title: "Властелин Колец", Author: "Толкин", Available: true},
		{Title: "Хоббит", Author: "Толкин", Available: true},
		{Title: "Сильмариллион", Author: "Толкин", Available: true}}

	for _, lib := range library {
		lib.Print()
	}
	library[0].Borrow()
	if !library[0].Borrow() {
		fmt.Println("Книга занята")
	}

	for _, lib := range library {
		lib.Print()
	}
	if library[0].Return() == true {
		fmt.Println("Книгу вернули")
	}

	for _, lib := range library {
		lib.Print()
	}
}
