package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

func filterByYear(books []Book, minyear int) []Book {
	result := []Book{}
	for _, b := range books {
		if b.Year >= minyear {
			result = append(result, b)
		}

	}
	return result

}

func main() {
	books := []Book{
		{Title: "Властелин колец", Author: "Толкин", Year: 2000},
		{Title: "Хоббит", Author: "Толкин", Year: 2005},
		{Title: "Сильмариллион", Author: "Толкин", Year: 2010},
		{Title: "Голлум", Author: "Толкин", Year: 2015}}

	book := filterByYear((books), 2010)

	for _, b := range book {
		fmt.Println(b.Title, " ", b.Author, b.Year)
	}

	data, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}
	fmt.Println(string(data))
	err = os.WriteFile("Books1", data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи: ", err)
		return
	}
	fmt.Print("Saved")
}
