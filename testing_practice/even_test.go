package main

import "testing"

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}

	return false
}
func TestIsEven(t *testing.T) {
	if !isEven(4) {
		t.Errorf("isEven(4)should be true")
	}
	if isEven(7) {
		t.Errorf("isEven(7) should be false")
	}

}
