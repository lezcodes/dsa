package quick_sort

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSortHelper(arr, low, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuickSortRandomPivot(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	quickSortRandomHelper(result, 0, len(result)-1)
	return result
}

func quickSortRandomHelper(arr []int, low, high int) {
	if low < high {
		randomIndex := low + rand.Intn(high-low+1)
		arr[randomIndex], arr[high] = arr[high], arr[randomIndex]

		pivotIndex := partition(arr, low, high)
		quickSortRandomHelper(arr, low, pivotIndex-1)
		quickSortRandomHelper(arr, pivotIndex+1, high)
	}
}

func QuickSortMedianOfThree(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	quickSortMedianHelper(result, 0, len(result)-1)
	return result
}

func quickSortMedianHelper(arr []int, low, high int) {
	if low < high {
		medianIndex := medianOfThree(arr, low, high)
		arr[medianIndex], arr[high] = arr[high], arr[medianIndex]

		pivotIndex := partition(arr, low, high)
		quickSortMedianHelper(arr, low, pivotIndex-1)
		quickSortMedianHelper(arr, pivotIndex+1, high)
	}
}

func medianOfThree(arr []int, low, high int) int {
	mid := (low + high) / 2

	if arr[low] > arr[mid] {
		if arr[mid] > arr[high] {
			return mid
		} else if arr[low] > arr[high] {
			return high
		} else {
			return low
		}
	} else {
		if arr[low] > arr[high] {
			return low
		} else if arr[mid] > arr[high] {
			return high
		} else {
			return mid
		}
	}
}

func QuickSortInPlace(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func Run() any {
	testCases := [][]int{
		{64, 34, 25, 12, 22, 11, 90},
		{5, 2, 8, 1, 9},
		{1},
		{},
		{3, 3, 3, 3},
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5},
		{10, 7, 8, 9, 1, 5},
	}

	results := make(map[string]any)

	for i, testCase := range testCases {
		original := make([]int, len(testCase))
		copy(original, testCase)

		basicSorted := QuickSort(testCase)
		randomSorted := QuickSortRandomPivot(testCase)
		medianSorted := QuickSortMedianOfThree(testCase)

		inPlaceTest := make([]int, len(testCase))
		copy(inPlaceTest, testCase)
		QuickSortInPlace(inPlaceTest)

		results[fmt.Sprintf("test_case_%d", i+1)] = map[string]any{
			"original":          original,
			"basic_quick_sort":  basicSorted,
			"random_pivot":      randomSorted,
			"median_of_three":   medianSorted,
			"in_place_sort":     inPlaceTest,
			"is_sorted_basic":   IsSorted(basicSorted),
			"is_sorted_random":  IsSorted(randomSorted),
			"is_sorted_median":  IsSorted(medianSorted),
			"is_sorted_inplace": IsSorted(inPlaceTest),
		}
	}

	results["algorithm_info"] = map[string]any{
		"name":             "Quick Sort",
		"time_complexity":  "O(n log n) average, O(n²) worst case",
		"space_complexity": "O(log n) average, O(n) worst case",
		"stable":           false,
		"in_place":         true,
		"variants_implemented": []string{
			"Basic (last element pivot)",
			"Random pivot",
			"Median-of-three pivot",
			"In-place sorting",
		},
	}

	return results
}
