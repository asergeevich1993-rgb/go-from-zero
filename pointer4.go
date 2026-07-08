package main

import "fmt"

type Counter struct {
	Value int
}

func reset(c *Counter) {
	c.Value = 0
}

func main() {

	c := Counter{Value: 100}
	reset(&c)
	fmt.Println(c)
	fmt.Println(c.Value)
}
