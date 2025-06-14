package bubble_sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Two elements - ascending",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Two elements - descending",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:     "With duplicates",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name:     "All same elements",
			input:    []int{7, 7, 7, 7, 7},
			expected: []int{7, 7, 7, 7, 7},
		},
		{
			name:     "Negative numbers",
			input:    []int{-3, -1, -4, -1, -5},
			expected: []int{-5, -4, -3, -1, -1},
		},
		{
			name:     "Mixed positive and negative",
			input:    []int{-3, 1, -4, 1, 5, -9, 2},
			expected: []int{-9, -4, -3, 1, 1, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v, expected %v", tt.input, result, tt.expected)
			}

			if len(tt.input) > 0 && reflect.DeepEqual(tt.input, result) && !reflect.DeepEqual(tt.input, tt.expected) {
				t.Error("BubbleSort should not modify the original array")
			}
		})
	}
}

func TestBubbleSortInPlace(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Random order",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			BubbleSortInPlace(input)

			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("BubbleSortInPlace(%v) = %v, expected %v", tt.input, input, tt.expected)
			}
		})
	}
}

func TestBubbleSortStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Empty array",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Single element",
			input:    []string{"hello"},
			expected: []string{"hello"},
		},
		{
			name:     "Alphabetical order",
			input:    []string{"banana", "apple", "cherry", "date"},
			expected: []string{"apple", "banana", "cherry", "date"},
		},
		{
			name:     "Already sorted",
			input:    []string{"apple", "banana", "cherry"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "With duplicates",
			input:    []string{"cat", "bat", "cat", "ant"},
			expected: []string{"ant", "bat", "cat", "cat"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSortStrings(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSortStrings(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestBubbleSortFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected []float64
	}{
		{
			name:     "Empty array",
			input:    []float64{},
			expected: []float64{},
		},
		{
			name:     "Single element",
			input:    []float64{3.14},
			expected: []float64{3.14},
		},
		{
			name:     "Multiple elements",
			input:    []float64{3.14, 2.71, 1.41, 0.57},
			expected: []float64{0.57, 1.41, 2.71, 3.14},
		},
		{
			name:     "With negative numbers",
			input:    []float64{-1.5, 2.7, -3.2, 0.0, 1.1},
			expected: []float64{-3.2, -1.5, 0.0, 1.1, 2.7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSortFloat64(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSortFloat64(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestBubbleSortWithSteps(t *testing.T) {
	input := []int{3, 1, 2}
	result, steps := BubbleSortWithSteps(input)

	expectedResult := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("BubbleSortWithSteps result = %v, expected %v", result, expectedResult)
	}

	if len(steps) < 2 {
		t.Errorf("Expected at least 2 steps, got %d", len(steps))
	}

	if !reflect.DeepEqual(steps[0], input) {
		t.Errorf("First step should be original array, got %v", steps[0])
	}

	if !reflect.DeepEqual(steps[len(steps)-1], expectedResult) {
		t.Errorf("Last step should be sorted array, got %v", steps[len(steps)-1])
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Error("Expected result to be a map")
	}

	requiredKeys := []string{"original_array", "sorted_array", "original_strings", "sorted_strings", "algorithm"}
	for _, key := range requiredKeys {
		if _, exists := resultMap[key]; !exists {
			t.Errorf("Expected result to contain '%s' key", key)
		}
	}
}

func BenchmarkBubbleSort10(b *testing.B) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for b.Loop() {
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSort100(b *testing.B) {
	arr := make([]int, 100)
	for i := range 100 {
		arr[i] = 100 - i
	}

	for b.Loop() {
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	arr := make([]int, 1000)
	for i := range 1000 {
		arr[i] = 1000 - i
	}

	for b.Loop() {
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSortInPlace10(b *testing.B) {
	originalArr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for b.Loop() {
		arr := make([]int, len(originalArr))
		copy(arr, originalArr)
		BubbleSortInPlace(arr)
	}
}

func BenchmarkBubbleSortInPlace100(b *testing.B) {
	originalArr := make([]int, 100)
	for i := range 100 {
		originalArr[i] = 100 - i
	}

	for b.Loop() {
		arr := make([]int, len(originalArr))
		copy(arr, originalArr)
		BubbleSortInPlace(arr)
	}
}

func BenchmarkBubbleSortBestCase(b *testing.B) {
	arr := make([]int, 100)
	for i := range 100 {
		arr[i] = i
	}

	for b.Loop() {
		BubbleSort(arr)
	}
}
