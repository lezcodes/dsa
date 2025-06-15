package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSortBasic(t *testing.T) {
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
		{[]int{38, 27, 43, 3, 9, 82, 10}, []int{3, 9, 10, 27, 38, 43, 82}},
	}

	for i, tc := range testCases {
		result := MergeSort(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, tc.expected, result)
		}
	}
}

func TestMergeSortBottomUp(t *testing.T) {
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
		result := MergeSortBottomUp(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, tc.expected, result)
		}
	}
}

func TestMergeSortInPlace(t *testing.T) {
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
		MergeSortInPlace(input)
		if !reflect.DeepEqual(input, tc.expected) {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, tc.expected, input)
		}
	}
}

func TestMergeSortStable(t *testing.T) {
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
		result := MergeSortStable(tc)
		if !IsSorted(result) {
			t.Errorf("Test case %d failed: result is not sorted %v", i+1, result)
		}
		if len(result) != len(tc) {
			t.Errorf("Test case %d failed: length mismatch", i+1)
		}
	}
}

func TestMergeSortOptimized(t *testing.T) {
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
		result := MergeSortOptimized(tc)
		if !IsSorted(result) {
			t.Errorf("Test case %d failed: result is not sorted %v", i+1, result)
		}
		if len(result) != len(tc) {
			t.Errorf("Test case %d failed: length mismatch", i+1)
		}
	}
}

func TestMerge(t *testing.T) {
	arr := []int{1, 3, 5, 2, 4, 6}
	merge(arr, 0, 2, 5)
	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Merge failed: expected %v, got %v", expected, arr)
	}
}

func TestMergeWithDuplicates(t *testing.T) {
	arr := []int{1, 3, 3, 2, 3, 4}
	merge(arr, 0, 2, 5)
	expected := []int{1, 2, 3, 3, 3, 4}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Merge with duplicates failed: expected %v, got %v", expected, arr)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 2, 8, 1, 9}
	insertionSort(arr, 0, 4)
	expected := []int{1, 2, 5, 8, 9}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Insertion sort failed: expected %v, got %v", expected, arr)
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

	result := MergeSort(arr)
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

	result := MergeSort(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNegativeNumbers(t *testing.T) {
	arr := []int{-5, 3, -1, 0, 8, -3}
	expected := []int{-5, -3, -1, 0, 3, 8}

	result := MergeSort(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestStabilityPreservation(t *testing.T) {
	type Item struct {
		value int
		index int
	}

	items := []Item{{3, 0}, {1, 1}, {3, 2}, {2, 3}}

	values := make([]int, len(items))
	for i, item := range items {
		values[i] = item.value
	}

	sortedValues := MergeSortStable(values)
	expected := []int{1, 2, 3, 3}

	if !reflect.DeepEqual(sortedValues, expected) {
		t.Errorf("Stability test failed: expected %v, got %v", expected, sortedValues)
	}
}

func TestEmptyAndSingleElement(t *testing.T) {
	empty := []int{}
	single := []int{42}

	emptyResult := MergeSort(empty)
	singleResult := MergeSort(single)

	if len(emptyResult) != 0 {
		t.Error("Empty array should remain empty")
	}

	if len(singleResult) != 1 || singleResult[0] != 42 {
		t.Error("Single element array should remain unchanged")
	}
}

func TestAllVariantsConsistency(t *testing.T) {
	testArray := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42}

	topDown := MergeSort(testArray)
	bottomUp := MergeSortBottomUp(testArray)
	stable := MergeSortStable(testArray)
	optimized := MergeSortOptimized(testArray)

	inPlace := make([]int, len(testArray))
	copy(inPlace, testArray)
	MergeSortInPlace(inPlace)

	if !reflect.DeepEqual(topDown, bottomUp) {
		t.Error("Top-down and bottom-up results differ")
	}

	if !reflect.DeepEqual(topDown, stable) {
		t.Error("Top-down and stable results differ")
	}

	if !reflect.DeepEqual(topDown, optimized) {
		t.Error("Top-down and optimized results differ")
	}

	if !reflect.DeepEqual(topDown, inPlace) {
		t.Error("Top-down and in-place results differ")
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

func BenchmarkMergeSortSmall(b *testing.B) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	for b.Loop() {
		MergeSort(arr)
	}
}

func BenchmarkMergeSortMedium(b *testing.B) {
	arr := make([]int, 100)
	for i := range 100 {
		arr[i] = 100 - i
	}

	for b.Loop() {
		MergeSort(arr)
	}
}

func BenchmarkMergeSortLarge(b *testing.B) {
	arr := make([]int, 1000)
	for i := range 1000 {
		arr[i] = 1000 - i
	}

	for b.Loop() {
		MergeSort(arr)
	}
}

func BenchmarkMergeSortBottomUp(b *testing.B) {
	arr := make([]int, 1000)
	for i := range 1000 {
		arr[i] = 1000 - i
	}

	for b.Loop() {
		MergeSortBottomUp(arr)
	}
}

func BenchmarkMergeSortOptimized(b *testing.B) {
	arr := make([]int, 1000)
	for i := range 1000 {
		arr[i] = 1000 - i
	}

	for b.Loop() {
		MergeSortOptimized(arr)
	}
}

func BenchmarkMergeSortInPlace(b *testing.B) {
	originalArr := make([]int, 1000)
	for i := range 1000 {
		originalArr[i] = 1000 - i
	}

	for b.Loop() {
		arr := make([]int, len(originalArr))
		copy(arr, originalArr)
		MergeSortInPlace(arr)
	}
}
