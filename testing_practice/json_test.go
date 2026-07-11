package main

import (
	"encoding/json"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func toJSON(p Person) ([]byte, error) {
	data, err := json.Marshal(p)

	return data, err
}

func fromJSON(data []byte) (Person, error) {

	var p Person
	err := json.Unmarshal(data, &p)

	return p, err
}

func TestJSON(t *testing.T) {
	person := []struct {
		input Person
		want  Person
	}{{Person{"Артур", 33}, Person{"Артур", 33}},
		{Person{"Мария", 23}, Person{"Мария", 23}},
		{Person{"", 0}, Person{"", 0}}}

	for _, p := range person {
		got, err := toJSON(p.input)

		persons, err := fromJSON(got)

		if err != nil {
			t.Errorf("fromJSON error: %v", err)
		}
		if persons != p.want {
			t.Errorf("got %v,want %v", persons, p.want)
		}
	}

}
