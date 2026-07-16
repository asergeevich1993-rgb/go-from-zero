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
		ch1 <- "Привет"
		ch1 <- "Пока"
	}()
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "Hello"
		ch2 <- "Bye"
	}()

	for i := 0; i < 4; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

}
