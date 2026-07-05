package main

import "fmt"

type Counter struct {
	Value int
}

func (c Counter) isPositive() bool {
	if c.Value > 0 {
		return true
	}
	return false
}

func main() {
	counter1 := Counter{Value: 5}
	counter2 := Counter{Value: -3}

	fmt.Println(counter1.Value, counter1.isPositive())
	fmt.Println(counter2.Value, counter2.isPositive())
}
