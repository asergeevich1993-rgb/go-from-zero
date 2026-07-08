package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string
	Pages  int
	Rating float64
}

func main() {
	book := []Book{
		{Title: "Властелин колец", Pages: 1000, Rating: 7.0},
		{Title: "Хоббит", Pages: 300, Rating: 8.0},
		{Title: "Игра престолов", Pages: 100500, Rating: 8.0},
	}

	data, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}
	fmt.Println(string(data))

	err = os.WriteFile("books.json", data, 0644)
	if err != nil {
		fmt.Print("Ошибка записи: ", err)
		return
	}
	fmt.Println("Saved")

	jdata, err1 := os.ReadFile("books.json")
	if err1 != nil {
		fmt.Println("Ошибка чтения: ", err)
		return
	}
	var books []Book
	err1 = json.Unmarshal(jdata, &books)
	if err != nil {
		fmt.Println("Ошибка десериализации: ", err1)
		return
	}
	for _, b := range books {
		fmt.Println(b.Title, " ", b.Pages, " ", b.Rating)
	}
}
