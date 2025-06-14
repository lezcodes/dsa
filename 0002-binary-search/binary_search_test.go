package binary_search

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Error("Expected result to be a map")
	}

	if resultMap["found"] != true {
		t.Error("Expected to find the target element")
	}

	if resultMap["index"] != 3 {
		t.Errorf("Expected index 3, got %v", resultMap["index"])
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

func TestSearchString(t *testing.T) {
	tests := []struct {
		name     string
		arr      []string
		target   string
		expected int
	}{
		{
			name:     "String found",
			arr:      []string{"apple", "banana", "cherry", "date", "elderberry"},
			target:   "cherry",
			expected: 2,
		},
		{
			name:     "String not found - before range",
			arr:      []string{"banana", "cherry", "date", "elderberry"},
			target:   "apple",
			expected: -1,
		},
		{
			name:     "String not found - after range",
			arr:      []string{"apple", "banana", "cherry", "date"},
			target:   "grape",
			expected: -1,
		},
		{
			name:     "String not found - in between",
			arr:      []string{"apple", "cherry", "elderberry"},
			target:   "banana",
			expected: -1,
		},
		{
			name:     "Empty string array",
			arr:      []string{},
			target:   "test",
			expected: -1,
		},
		{
			name:     "Single string found",
			arr:      []string{"test"},
			target:   "test",
			expected: 0,
		},
		{
			name:     "Single string not found",
			arr:      []string{"test"},
			target:   "other",
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchString(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("SearchString(%v, %s) = %d; expected %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func TestSearchFloat64(t *testing.T) {
	tests := []struct {
		name     string
		arr      []float64
		target   float64
		expected int
	}{
		{
			name:     "Float found",
			arr:      []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			target:   3.3,
			expected: 2,
		},
		{
			name:     "Float not found",
			arr:      []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			target:   6.6,
			expected: -1,
		},
		{
			name:     "Float found at beginning",
			arr:      []float64{0.5, 1.5, 2.5, 3.5},
			target:   0.5,
			expected: 0,
		},
		{
			name:     "Float found at end",
			arr:      []float64{0.5, 1.5, 2.5, 3.5},
			target:   3.5,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchFloat64(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("SearchFloat64(%v, %f) = %d; expected %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func TestSearchGeneric(t *testing.T) {
	intLess := func(a, b int) bool { return a < b }
	intResult := SearchGeneric([]int{1, 2, 3, 4, 5}, 3, intLess)
	if intResult != 2 {
		t.Errorf("SearchGeneric with ints failed: got %d, expected 2", intResult)
	}

	stringLess := func(a, b string) bool { return a < b }
	stringResult := SearchGeneric([]string{"a", "b", "c", "d"}, "b", stringLess)
	if stringResult != 1 {
		t.Errorf("SearchGeneric with strings failed: got %d, expected 1", stringResult)
	}

	notFoundResult := SearchGeneric([]int{1, 2, 3}, 5, intLess)
	if notFoundResult != -1 {
		t.Errorf("SearchGeneric not found case failed: got %d, expected -1", notFoundResult)
	}

	emptyResult := SearchGeneric([]int{}, 5, intLess)
	if emptyResult != -1 {
		t.Errorf("SearchGeneric empty array case failed: got %d, expected -1", emptyResult)
	}
}

func BenchmarkSearchInt_Small(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	target := 3

	for b.Loop() {
		SearchInt(arr, target)
	}
}

func BenchmarkSearchInt_Medium(b *testing.B) {
	arr := make([]int, 1000)
	for i := range 1000 {
		arr[i] = i
	}
	target := 500

	for b.Loop() {
		SearchInt(arr, target)
	}
}

func BenchmarkSearchInt_Large(b *testing.B) {
	arr := make([]int, 100000)
	for i := range 100000 {
		arr[i] = i
	}
	target := 50000

	for b.Loop() {
		SearchInt(arr, target)
	}
}

func BenchmarkSearchInt_VeryLarge(b *testing.B) {
	arr := make([]int, 1000000)
	for i := range 1000000 {
		arr[i] = i
	}
	target := 500000

	for b.Loop() {
		SearchInt(arr, target)
	}
}

func BenchmarkSearchString(b *testing.B) {
	arr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	target := "cherry"

	for b.Loop() {
		SearchString(arr, target)
	}
}

func BenchmarkSearchGeneric(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	less := func(a, b int) bool { return a < b }

	for b.Loop() {
		SearchGeneric(arr, target, less)
	}
}

func TestRunOutput(t *testing.T) {
	result := Run()
	resultMap := result.(map[string]any)

	expectedArray := []int{11, 12, 22, 25, 34, 64, 90}
	actualArray := resultMap["array"].([]int)

	if !reflect.DeepEqual(actualArray, expectedArray) {
		t.Errorf("Expected array %v, got %v", expectedArray, actualArray)
	}

	if resultMap["target"] != 25 {
		t.Errorf("Expected target 25, got %v", resultMap["target"])
	}

	if resultMap["index"] != 3 {
		t.Errorf("Expected index 3, got %v", resultMap["index"])
	}

	if resultMap["found"] != true {
		t.Errorf("Expected found true, got %v", resultMap["found"])
	}
}
