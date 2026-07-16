package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
}

func (p Person) SayHi() {
	fmt.Print("привет я,", p.Name, "!")
}
func main() {
	p := Person{Name: "Артур"}
	go p.SayHi()
	time.Sleep(100 * time.Millisecond)
}
