package main

import "fmt"

func main() {

	var num, rows, reps int
	fmt.Println("Число: ")
	fmt.Println("строк: ")
	fmt.Println("повторы: ")
	fmt.Scan(&num, &rows, &reps)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= reps; j++ {
			fmt.Print(num)
		}
		fmt.Println("")
	}

}
