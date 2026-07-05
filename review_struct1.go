package main

import "fmt"

type Movie struct {
	Title  string
	Year   int
	Rating float64
}

func BestMovie(movies []Movie) Movie {

	best := movies[0]
	for _, m := range movies {
		if m.Rating > best.Rating {
			best = m
		}

	}
	return best

}

func main() {
	movie := []Movie{{Title: "Властелин Колец", Year: 2001, Rating: 10.0},
		{Title: "Хоббит", Year: 2013, Rating: 9.0},
		{Title: "Голлум", Year: 2027, Rating: 6.0}}
	for _, m := range movie {
		fmt.Println(m.Title, ":", m.Year, " ", m.Rating)
	}
	winner := BestMovie(movie)
	fmt.Println("Лучший: ", winner.Title, ":", winner.Rating)
}
