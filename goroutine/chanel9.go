package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case num := <-ch:
		fmt.Println("получили: ", num)
	default:
		fmt.Println("Канал пуст")
	}
}
