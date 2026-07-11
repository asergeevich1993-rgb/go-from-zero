package main

import (
	"fmt"
	"testing"
)

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на 0")
	}
	return a / b, nil
}

func TestSafeDivide(t *testing.T) {
	nums := []struct {
		a, b    int
		Want    int
		ErrWant bool
	}{
		{4, 2, 2, false},
		{4, 0, 0, true},
		{20, 2, 10, false},
		{10, 5, 2, false}}
	for _, n := range nums {
		got, err := safeDivide(n.a, n.b)
		if n.ErrWant {
			if err == nil {
				t.Errorf("Должна была быть ошибка")
			}
		} else {
			if err != nil {
				t.Errorf("непредвиденная ошибка")
			}
			if got != n.Want {
				t.Errorf("%d/%d=%v,want %d", n.a, n.b, got, n.Want)
			}
		}

	}

}
