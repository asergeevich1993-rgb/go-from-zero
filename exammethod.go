package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Inc() {
	c.Value += 1

}

func (c *Counter) Dec() {
	c.Value -= 1

}

func (c Counter) Show() {
	fmt.Println(c.Value)

}
func main() {
	count := Counter{Value: 10}
	count.Inc()
	count.Dec()
	count.Inc()
	count.Show()
}
