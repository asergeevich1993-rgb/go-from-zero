package main

import "fmt"

type Checker interface {
	Check() bool
}

type EvenNumber struct {
	Value int
}

func (nums EvenNumber) Check() bool {
	return nums.Value%2 == 0

}

type PositiveNumber struct {
	Value int
}

func (nums PositiveNumber) Check() bool {
	return nums.Value > 0
}
func filter(items []Checker) []Checker {
	result := []Checker{}
	for _, i := range items {
		if i.Check() == true {
			result = append(result, i)
		}
	}
	return result
}

func main() {

	items := []Checker{EvenNumber{Value: 2}, EvenNumber{Value: 5}, PositiveNumber{Value: -3}, PositiveNumber{Value: 3}}

	result := filter(items)
	fmt.Println(result)

}
