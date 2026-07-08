package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name   string
	Grades []int
}

func main() {
	student := []Student{{Name: "Артур", Grades: []int{4, 3, 5, 3, 5}},
		{Name: "Мария", Grades: []int{5, 4, 5, 3, 4}}}

	data, err := json.Marshal(student)
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}
	err = os.WriteFile("grades.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}
	fmt.Println("Saved")

	jData, err1 := os.ReadFile("grades.json")
	if err1 != nil {
		fmt.Println("Ошибка чтения: ", err)
		return
	}
	var students []Student
	err1 = json.Unmarshal(jData, &students)
	if err1 != nil {
		fmt.Println("Ошибка десериализации: ", err1)
		return
	}
	for _, s := range students {
		fmt.Println(s.Name, ":", s.Grades)
	}

}
