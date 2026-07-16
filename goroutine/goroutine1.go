package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	go say("Привет")
	say("Пока")
}
