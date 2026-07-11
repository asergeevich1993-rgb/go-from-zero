package main

import (
	"fmt"
	"testing"
)

func safeDivide1(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на 0")
	}
	return a / b, nil

}

func TestSafeDivide1(t *testing.T) {
	result := []struct {
		a, b    int
		want    int
		errwant bool
	}{
		{4, 2, 2, false},
		{8, 4, 2, false},
		{10, 0, 0, true}}

	for _, r := range result {
		got, err := safeDivide1(r.a, r.b)
		if r.errwant {
			if err == nil {
				t.Errorf("Ожидалась ошибка")
			}
		} else {
			if err != nil {
				t.Errorf("непредвиденная ошибка")
			}
			if got != r.want {
				t.Errorf("%d/%d = %v, want %d", r.a, r.b, got, r.want)
			}
		}

	}
}
