package fibonacci

import (
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result from Run()")
	}
}

func TestFibonacciIterative(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range tests {
		result := FibonacciIterative(test.n)
		if result != test.expected {
			t.Errorf("For n=%d, expected %d, got %d", test.n, test.expected, result)
		}
	}
}

func TestFibonacciRecursive(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range tests {
		result := FibonacciRecursive(test.n)
		if result != test.expected {
			t.Errorf("For n=%d, expected %d, got %d", test.n, test.expected, result)
		}
	}
}
