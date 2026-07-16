package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Println("привет,", name, "!")
}

func main() {

	go sayHello("Артур")
	time.Sleep(100 * time.Millisecond)
}
