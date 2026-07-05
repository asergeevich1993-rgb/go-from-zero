package main

import "fmt"

type City struct {
	Name       string
	Population float64
	Area       float64
}

func mostPopulation(peoples []City) City {
	most := peoples[0]

	for _, p := range peoples {
		if p.Population/p.Area > most.Population/most.Area {
			most = p
		}
	}
	return most
}
func main() {

	Cities := []City{
		{Name: "Москва", Population: 100, Area: 150.55},
		{Name: "Владивосток", Population: 40, Area: 90.53},
		{Name: "Питер", Population: 90, Area: 130.51},
		{Name: "Деревня", Population: 10, Area: 10.5}}

	for _, c := range Cities {
		fmt.Println(c.Name, ":", c.Population, " ", c.Area)
	}
	Most := mostPopulation(Cities)
	fmt.Println(Most.Name, ":", Most.Population, " ", Most.Area)

}
