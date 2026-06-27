package main

import (
	"fmt"
	"strings"
)

func filterByKey(m map[string]int, letters string) map[string]int {
	result := map[string]int{}
	for key, value := range m {
		if strings.HasPrefix(key, "к") {
			result[key] = value

		}

	}
	return result
}

func main() {

	result := filterByKey(map[string]int{"ручки": 100, "карандаши": 200, "тетради": 50}, "к")
	fmt.Print(result)

}
