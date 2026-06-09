package main

import "fmt"

func main() {
	var temperature int
	fmt.Print("Введи температуру или иди нах: ")
	fmt.Scan(&temperature)

	if temperature < -20 {
		fmt.Print("Дубак сиди дома")
	} else if temperature >= -20 && temperature <= -1 {
		fmt.Print("Дубак одень пуховик")
	} else if temperature >= 0 && temperature <= 9 {
		fmt.Print("Прохладно надень куртку")
	} else if temperature >= 10 && temperature <= 19 {
		fmt.Print("Одень толстовку")
	} else if temperature >= 20 && temperature <= 29 {
		fmt.Print("Одень футболку")
	} else if temperature >= 30 && temperature <= 60 {
		fmt.Print("Вали на пляж с спф кремом")
	} else {
		fmt.Print("настал конец света и солнце переживает свою последнюю фазу жизни")
	}

}
