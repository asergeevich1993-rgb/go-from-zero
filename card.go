package main

import "fmt"

func main() {

	var firstname string = "Artur"
	var lastname string = "Shevchenko"
	var age int = 33
	var height float64 = 1.85
	var isStudent bool = true

	fmt.Println("Name:", firstname)
	fmt.Println("Familia:", lastname)
	fmt.Println("Vozrast:", age)
	fmt.Println("Ves:", height)
	fmt.Print("Student: ", isStudent)
	fmt.Println("")

	fmt.Printf("Визитка:%s %s Возраст: %d , Рост: %.2f ,Студентота: %t\n", firstname, lastname, age, height, isStudent)

	isStudent = false
	fmt.Printf("Визитка:%s %s Возраст: %d , Рост: %.2f ,Студентота: %t\n", firstname, lastname, age, height, isStudent)

}
