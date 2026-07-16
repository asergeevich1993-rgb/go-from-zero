package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Первая"
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Вторая"
	}()
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			time.Sleep(50 * time.Millisecond)
			fmt.Println("Жду")
		}
	}
}
