package main

import "fmt"

type Azazatel struct {
	Value int
}

func newCounter(value int) *Azazatel {
	return &Azazatel{Value: value}
}
func main() {

	c := newCounter(50)
	fmt.Println(c)

}
