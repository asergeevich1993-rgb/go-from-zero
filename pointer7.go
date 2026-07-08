package main

import "fmt"

type Counter struct {
	Value int
}

func Increment(c *Counter) {
	c.Value += 1

}

func main() {
	c := Counter{Value: 10}
	Increment(&c)
	fmt.Println(c)
}
