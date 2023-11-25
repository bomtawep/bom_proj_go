package main

import "testing"

func TestAddition(t *testing.T) {
	tests := []struct {
		a, b, sum int
	}{
		{1, 2, 3},
		{4, 5, 9},
		{-1, 1, 0},
	}

	for _, test := range tests {
		result := add(test.a, test.b)
		if result != test.sum {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", result, test.sum)
		}
	}
}

func add(a int, b int) interface{} {
	return a + b
}
