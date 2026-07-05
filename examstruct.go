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
	movies := []Movie{{Title: "Властелин Колец", Year: 2001, Rating: 10.0}, {Title: "Хоббит", Year: 2013, Rating: 9.0},
		{Title: "Аватар", Year: 2009, Rating: 8.0}}

	for _, m := range movies {
		fmt.Println(m.Rating, ":", m.Title, " ", m.Year)
	}
	best := BestMovie(movies)

	fmt.Println(best.Rating, ":", best.Title)

}
