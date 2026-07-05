package main

import "fmt"

type Filter interface {
	Pass() bool
}

type Adult struct {
	Age int
}

func (a Adult) Pass() bool {
	return a.Age >= 18

}

type Rich struct {
	Money int
}

func (r Rich) Pass() bool {
	return r.Money >= 1000
}

func filter(items []Filter) []Filter {
	result := []Filter{}

	for _, i := range items {
		if i.Pass() == true {
			result = append(result, i)
		}
	}
	return result

}

func main() {
	result := []Filter{Adult{Age: 16}, Adult{Age: 20}, Rich{Money: 500}, Rich{Money: 5000}}

	result1 := filter(result)
	fmt.Println(result1)
}
