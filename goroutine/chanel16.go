package main

import "fmt"

func main() {

	ch := make(chan string)

	select {
	default:
		fmt.Print("no data")
	}

}
