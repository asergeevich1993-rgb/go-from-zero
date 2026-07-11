package main

import (
	"reflect"
	"testing"
)

func wordCount(words []string) map[string]int {
	counter := map[string]int{}
	for _, w := range words {
		counter[w]++

	}
	return counter
}

func TestWordCount(t *testing.T) {
	words := []struct {
		words []string
		want  map[string]int
	}{{[]string{"го", "го", "это"},
		map[string]int{"го": 2, "это": 1}},
		{[]string{"один"}, map[string]int{"один": 1}},
		{[]string{}, map[string]int{}}}

	for _, w := range words {
		got := wordCount(w.words)
		if !reflect.DeepEqual(got, w.want) {
			t.Errorf("%v не совпадает,want %v", got, w.want)

		}
	}

}
