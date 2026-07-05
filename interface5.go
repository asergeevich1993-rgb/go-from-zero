package main

import "fmt"

type Vehicle interface {
	Drive() string

	Stop() string
}

type Car struct {
	Start string
	stop  string
}

func (c Car) Drive() string {
	return "Машина " + c.Start

}
func (cc Car) Stop() string {
	return "Машина " + cc.stop
}

type Bike struct {
	Start string
	stop  string
}

func (b Bike) Drive() string {
	return "мотоцикл: " + b.Start
}
func (bb Bike) Stop() string {
	return "Мотоцикл: " + bb.stop
}

func testDrive(v Vehicle) {
	fmt.Println(v.Drive())
	fmt.Println(v.Stop())
}

func main() {
	m := Car{Start: "Еду", stop: "Остановка!"}
	b := Bike{Start: "Еду на двух колесах", stop: "Подножка!"}

	testDrive(m)
	testDrive(b)
}
