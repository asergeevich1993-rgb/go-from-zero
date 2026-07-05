package main

import "fmt"

type Vehicle interface {
	Drive() string

	Stop() string
}

type Car struct {
	CarStart string
	CarStop  string
}

func (c Car) Drive() string {
	return "Машина " + c.CarStart

}
func (cc Car) Stop() string {
	return "Машина " + cc.CarStop
}

type Bike struct {
	BikeStart string
	BikeStop  string
}

func (b Bike) Drive() string {
	return "мотоцикл: " + b.BikeStart
}
func (bb Bike) Stop() string {
	return "Мотоцикл: " + bb.BikeStop
}

func main() {
	V := []Vehicle{Car{CarStart: "Еду", CarStop: "Остановка!"},
		Bike{BikeStart: "Еду на двух колесах", BikeStop: "Подножка!"}}

	for _, s := range V {
		fmt.Println(s.Drive())
		fmt.Println(s.Stop())
	}
}
