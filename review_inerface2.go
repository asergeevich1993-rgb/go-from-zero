package main

import "fmt"

type Calculator interface {
	Add(amount int)
	Value() int
}

type Sum struct {
	Result int
}

func (s *Sum) Add(amount int) {
	s.Result = s.Result + amount
}

func (s Sum) Value() int {
	return s.Result
}

type Multiply struct {
	Result int
}

func (m *Multiply) Add(amount int) {
	m.Result = m.Result * amount
}
func (m Multiply) Value() int {
	return m.Result
}

func testCalc(c Calculator) {
	c.Add(3)
	c.Add(5)
	fmt.Println(c.Value())
}

func main() {
	sum := Sum{Result: 10}
	multiply := Multiply{Result: 2}

	testCalc(&sum)
	testCalc(&multiply)
}
