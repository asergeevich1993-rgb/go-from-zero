package main

import "fmt"

func inverMap(m map[string]int) map[int]string {

	result := map[int]string{}
	for key, value := range m {
		result[value] = key
	}
	return result
}

func main() {

	result := inverMap(map[string]int{"ручки": 100, "тетради": 50})
	fmt.Print(result)

}
