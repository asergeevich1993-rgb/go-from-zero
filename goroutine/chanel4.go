package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "первый"
	ch <- "второй"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
