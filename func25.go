package main

import "fmt"

func mergeMaps(a, b map[string]int) map[string]int {
	sum := map[string]int{}

	for key, value := range a {
		sum[key] = value
	}
	for key, value := range b {
		sum[key] = sum[key] + value
	}
	return sum

}

func main() {

	a := map[string]int{"ручки": 100, "тетради": 50}
	b := map[string]int{"ручки": 30, "карандаши": 200}
	resuls := mergeMaps(a, b)
	fmt.Print(resuls)

}
