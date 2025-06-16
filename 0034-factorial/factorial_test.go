package factorial

import (
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result from Run()")
	}
}

func TestFactorialIterative(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}

	for _, test := range tests {
		result := FactorialIterative(test.n)
		if result != test.expected {
			t.Errorf("For n=%d, expected %d, got %d", test.n, test.expected, result)
		}
	}
}

func TestFactorialRecursive(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}

	for _, test := range tests {
		result := FactorialRecursive(test.n)
		if result != test.expected {
			t.Errorf("For n=%d, expected %d, got %d", test.n, test.expected, result)
		}
	}
}
