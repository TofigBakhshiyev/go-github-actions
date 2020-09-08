package main

import (
	"fmt"
	"testing"
)

func TestPass(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{9, 34},
		{5, 5},
		{6, 8},
		{11, 89},
		{19, 4181},
	}

	for _, test := range tests {
		if output := Fibonacci(test.input); output != test.expected {
			t.Error("Test failed: output {}, expected {}", test.input, test.expected)
		}
	}
	fmt.Println("Done")
}
