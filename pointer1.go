package main

import "fmt"

func main() {
	x := 5
	p := &x

	fmt.Println("x:", x)   // 5
	fmt.Println("p:", p)   // адрес (0x...)
	fmt.Println("*p:", *p) // 5 (значение по адресу)

	*p = 10
	fmt.Println("x:", x) // 10 (оригинал изменился!)
}
