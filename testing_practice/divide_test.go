package main

import (
	"fmt"
	"testing"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на 0")
	}

	return a / b, nil
}
func TestDivide(t *testing.T) {
	tests := []struct {
		a, b    int
		want    int
		wantErr bool
	}{
		{6, 3, 2, false},
		{5, 2, 2, false},
		{5, 0, 0, true}}

	for _, tt := range tests {
		got, err := divide(tt.a, tt.b)

		if tt.wantErr {
			if err == nil {
				t.Errorf("Ожидалась ошибка")
			}
		} else {
			if err != nil {
				t.Errorf("Неожиданная ошибка %v", err)
			}

			if got != tt.want {
				t.Errorf("divide(%d,%d)=%d,want %d ", tt.a, tt.b, got, tt.want)
			}
		}
	}

}
