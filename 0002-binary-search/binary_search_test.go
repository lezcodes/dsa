package binary_search

import (
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result from Run()")
	}
}

func TestSearchInt(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Element found at beginning",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Element found at end",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "Element found in middle",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "Element not found - too small",
			arr:      []int{1, 2, 3, 4, 5},
			target:   0,
			expected: -1,
		},
		{
			name:     "Element not found - too large",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "Element not found - in between",
			arr:      []int{1, 3, 5, 7, 9},
			target:   4,
			expected: -1,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "Single element found",
			arr:      []int{42},
			target:   42,
			expected: 0,
		},
		{
			name:     "Single element not found",
			arr:      []int{42},
			target:   1,
			expected: -1,
		},
		{
			name:     "Large sorted array",
			arr:      []int{1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50},
			target:   25,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchInt(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("SearchInt(%v, %d) = %d; expected %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
