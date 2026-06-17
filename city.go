package main

import "fmt"

func main() {

	city := []string{"Владивосток", "Питер", "Москва", "Калининград"}
	var choise int
	fmt.Println(len(city))
	fmt.Println(city[0], city[3])
	fmt.Println("из какого вы города? Владивосток (1), Питер(2), Москва(3), Калининград(4)")
	fmt.Scan(&choise)
	for i := 0; i < len(city); i++ {

	}
	if choise == 1 {
		fmt.Print("Красавчик")
	} else if choise == 2 {
		fmt.Print("ну тоже норм")
	} else if choise == 3 {
		fmt.Print("фу, гей")
	} else {
		fmt.Print("деревня")
	}

}
