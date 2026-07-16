package main

import (
	"fmt"
	"time"
)

func main() {

	words1 := make(chan string, 1)
	words2 := make(chan string, 1)
	go func() {
		time.Sleep(150 * time.Millisecond)
		words1 <- "Привет"
	}()
	go func() {
		time.Sleep(300 * time.Millisecond)
		words2 <- "пока"
	}()

	for i := 1; i <= 4; i++ {
		select {
		case msg1 := <-words1:
			fmt.Println(msg1)
		case msg2 := <-words2:
			fmt.Println(msg2)
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Ожидание...")
		}

	}
	fmt.Println("Finish")

}
