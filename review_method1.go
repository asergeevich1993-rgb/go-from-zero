package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) inc() {
	c.Value += 1

}

func (c *Counter) Dec() {
	c.Value -= 1
}

func (c Counter) Show() {
	fmt.Println(c.Value)
}
func main() {
	counter := Counter{Value: 10}

	counter.Show()
	counter.inc()
	counter.inc()
	counter.Show()
	counter.Dec()
	counter.Dec()
	counter.Show()

}
