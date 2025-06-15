package quick_sort

import (
	"reflect"
	"testing"
)

func TestQuickSortBasic(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
		{[]int{5, 2, 8, 1, 9}, []int{1, 2, 5, 8, 9}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{3, 3, 3, 3}, []int{3, 3, 3, 3}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for i, tc := range testCases {
		result := QuickSort(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, tc.expected, result)
		}
	}
}

func TestQuickSortRandomPivot(t *testing.T) {
	testCases := [][]int{
		{64, 34, 25, 12, 22, 11, 90},
		{5, 2, 8, 1, 9},
		{1},
		{},
		{3, 3, 3, 3},
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5},
	}

	for i, tc := range testCases {
		result := QuickSortRandomPivot(tc)
		if !IsSorted(result) {
			t.Errorf("Test case %d failed: result is not sorted %v", i+1, result)
		}
		if len(result) != len(tc) {
			t.Errorf("Test case %d failed: length mismatch", i+1)
		}
	}
}

func TestQuickSortMedianOfThree(t *testing.T) {
	testCases := [][]int{
		{64, 34, 25, 12, 22, 11, 90},
		{5, 2, 8, 1, 9},
		{1},
		{},
		{3, 3, 3, 3},
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5},
	}

	for i, tc := range testCases {
		result := QuickSortMedianOfThree(tc)
		if !IsSorted(result) {
			t.Errorf("Test case %d failed: result is not sorted %v", i+1, result)
		}
		if len(result) != len(tc) {
			t.Errorf("Test case %d failed: length mismatch", i+1)
		}
	}
}

func TestQuickSortInPlace(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
		{[]int{5, 2, 8, 1, 9}, []int{1, 2, 5, 8, 9}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{3, 3, 3, 3}, []int{3, 3, 3, 3}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for i, tc := range testCases {
		input := make([]int, len(tc.input))
		copy(input, tc.input)
		QuickSortInPlace(input)
		if !reflect.DeepEqual(input, tc.expected) {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, tc.expected, input)
		}
	}
}

func TestPartition(t *testing.T) {
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	pivotIndex := partition(arr, 0, len(arr)-1)

	if pivotIndex < 0 || pivotIndex >= len(arr) {
		t.Errorf("Invalid pivot index: %d", pivotIndex)
	}

	pivot := arr[pivotIndex]
	for i := range pivotIndex {
		if arr[i] > pivot {
			t.Errorf("Element %d at index %d is greater than pivot %d", arr[i], i, pivot)
		}
	}

	for i := pivotIndex + 1; i < len(arr); i++ {
		if arr[i] < pivot {
			t.Errorf("Element %d at index %d is less than pivot %d", arr[i], i, pivot)
		}
	}
}

func TestMedianOfThree(t *testing.T) {
	testCases := []struct {
		arr      []int
		low      int
		high     int
		expected int
	}{
		{[]int{3, 1, 2}, 0, 2, 2},
		{[]int{1, 2, 3}, 0, 2, 1},
		{[]int{2, 3, 1}, 0, 2, 0},
		{[]int{5, 1, 3, 9, 2}, 0, 4, 2},
	}

	for i, tc := range testCases {
		result := medianOfThree(tc.arr, tc.low, tc.high)
		if result != tc.expected {
			t.Errorf("Test case %d failed: expected index %d, got %d", i+1, tc.expected, result)
		}
	}
}

func TestIsSorted(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{1}, true},
		{[]int{}, true},
		{[]int{2, 2, 2, 2}, true},
		{[]int{1, 2, 2, 3, 3}, true},
	}

	for i, tc := range testCases {
		result := IsSorted(tc.arr)
		if result != tc.expected {
			t.Errorf("Test case %d failed: expected %v, got %v for array %v", i+1, tc.expected, result, tc.arr)
		}
	}
}

func TestLargeArray(t *testing.T) {
	size := 1000
	arr := make([]int, size)
	for i := range size {
		arr[i] = size - i
	}

	result := QuickSort(arr)
	if !IsSorted(result) {
		t.Error("Large array not sorted correctly")
	}

	if len(result) != size {
		t.Errorf("Expected length %d, got %d", size, len(result))
	}

	for i := range size {
		if result[i] != i+1 {
			t.Errorf("Expected %d at index %d, got %d", i+1, i, result[i])
			break
		}
	}
}

func TestDuplicateElements(t *testing.T) {
	arr := []int{5, 2, 8, 2, 9, 1, 5, 5}
	expected := []int{1, 2, 2, 5, 5, 5, 8, 9}

	result := QuickSort(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNegativeNumbers(t *testing.T) {
	arr := []int{-5, 3, -1, 0, 8, -3}
	expected := []int{-5, -3, -1, 0, 3, 8}

	result := QuickSort(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	data, ok := result.(map[string]any)
	if !ok {
		t.Error("Expected map result")
	}

	if _, exists := data["algorithm_info"]; !exists {
		t.Error("Expected algorithm_info in result")
	}

	if _, exists := data["test_case_1"]; !exists {
		t.Error("Expected test cases in result")
	}
}

func BenchmarkQuickSortSmall(b *testing.B) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	for b.Loop() {
		QuickSort(arr)
	}
}

func BenchmarkQuickSortMedium(b *testing.B) {
	arr := make([]int, 100)
	for i := range 100 {
		arr[i] = 100 - i
	}

	for b.Loop() {
		QuickSort(arr)
	}
}

func BenchmarkQuickSortLarge(b *testing.B) {
	arr := make([]int, 1000)
	for i := range 1000 {
		arr[i] = 1000 - i
	}

	for b.Loop() {
		QuickSort(arr)
	}
}

func BenchmarkQuickSortRandomPivot(b *testing.B) {
	arr := make([]int, 1000)
	for i := range 1000 {
		arr[i] = 1000 - i
	}

	for b.Loop() {
		QuickSortRandomPivot(arr)
	}
}

func BenchmarkQuickSortMedianOfThree(b *testing.B) {
	arr := make([]int, 1000)
	for i := range 1000 {
		arr[i] = 1000 - i
	}

	for b.Loop() {
		QuickSortMedianOfThree(arr)
	}
}

func BenchmarkQuickSortInPlace(b *testing.B) {
	originalArr := make([]int, 1000)
	for i := range 1000 {
		originalArr[i] = 1000 - i
	}

	for b.Loop() {
		arr := make([]int, len(originalArr))
		copy(arr, originalArr)
		QuickSortInPlace(arr)
	}
}
