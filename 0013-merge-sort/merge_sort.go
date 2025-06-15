package merge_sort

import (
	"fmt"
)

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	mergeSortHelper(result, 0, len(result)-1)
	return result
}

func mergeSortHelper(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2

		mergeSortHelper(arr, left, mid)
		mergeSortHelper(arr, mid+1, right)

		merge(arr, left, mid, right)
	}
}

func merge(arr []int, left, mid, right int) {
	leftSize := mid - left + 1
	rightSize := right - mid

	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)

	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])

	i, j, k := 0, 0, left

	for i < leftSize && j < rightSize {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < leftSize {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < rightSize {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

func MergeSortBottomUp(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)

	n := len(result)
	for size := 1; size < n; size *= 2 {
		for left := 0; left < n-size; left += 2 * size {
			mid := left + size - 1
			right := min(left+2*size-1, n-1)
			merge(result, left, mid, right)
		}
	}

	return result
}

func MergeSortInPlace(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mergeSortHelper(arr, 0, len(arr)-1)
}

func MergeSortStable(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	mergeSortStableHelper(result, 0, len(result)-1)
	return result
}

func mergeSortStableHelper(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2

		mergeSortStableHelper(arr, left, mid)
		mergeSortStableHelper(arr, mid+1, right)

		mergeStable(arr, left, mid, right)
	}
}

func mergeStable(arr []int, left, mid, right int) {
	leftSize := mid - left + 1
	rightSize := right - mid

	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)

	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])

	i, j, k := 0, 0, left

	for i < leftSize && j < rightSize {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < leftSize {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < rightSize {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

func MergeSortOptimized(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	temp := make([]int, len(result))
	mergeSortOptimizedHelper(result, temp, 0, len(result)-1)
	return result
}

func mergeSortOptimizedHelper(arr, temp []int, left, right int) {
	if left < right {
		if right-left <= 10 {
			insertionSort(arr, left, right)
			return
		}

		mid := left + (right-left)/2

		mergeSortOptimizedHelper(arr, temp, left, mid)
		mergeSortOptimizedHelper(arr, temp, mid+1, right)

		if arr[mid] <= arr[mid+1] {
			return
		}

		mergeOptimized(arr, temp, left, mid, right)
	}
}

func mergeOptimized(arr, temp []int, left, mid, right int) {
	copy(temp[left:right+1], arr[left:right+1])

	i, j, k := left, mid+1, left

	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}

	for i <= mid {
		arr[k] = temp[i]
		i++
		k++
	}
}

func insertionSort(arr []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
		{38, 27, 43, 3, 9, 82, 10},
	}

	results := make(map[string]any)

	for i, testCase := range testCases {
		original := make([]int, len(testCase))
		copy(original, testCase)

		topDownSorted := MergeSort(testCase)
		bottomUpSorted := MergeSortBottomUp(testCase)
		stableSorted := MergeSortStable(testCase)
		optimizedSorted := MergeSortOptimized(testCase)

		inPlaceTest := make([]int, len(testCase))
		copy(inPlaceTest, testCase)
		MergeSortInPlace(inPlaceTest)

		results[fmt.Sprintf("test_case_%d", i+1)] = map[string]any{
			"original":            original,
			"top_down_sort":       topDownSorted,
			"bottom_up_sort":      bottomUpSorted,
			"stable_sort":         stableSorted,
			"optimized_sort":      optimizedSorted,
			"in_place_sort":       inPlaceTest,
			"is_sorted_topdown":   IsSorted(topDownSorted),
			"is_sorted_bottomup":  IsSorted(bottomUpSorted),
			"is_sorted_stable":    IsSorted(stableSorted),
			"is_sorted_optimized": IsSorted(optimizedSorted),
			"is_sorted_inplace":   IsSorted(inPlaceTest),
		}
	}

	results["algorithm_info"] = map[string]any{
		"name":             "Merge Sort",
		"time_complexity":  "O(n log n) - guaranteed",
		"space_complexity": "O(n) - for temporary arrays",
		"stable":           true,
		"in_place":         false,
		"variants_implemented": []string{
			"Top-down (recursive)",
			"Bottom-up (iterative)",
			"Stable merge sort",
			"Optimized (with insertion sort for small arrays)",
			"In-place sorting",
		},
		"advantages": []string{
			"Guaranteed O(n log n) time complexity",
			"Stable sorting algorithm",
			"Predictable performance",
			"Good for large datasets",
			"Parallelizable",
		},
		"disadvantages": []string{
			"Requires O(n) extra space",
			"Not in-place",
			"Slower than quicksort in practice for small arrays",
		},
	}

	return results
}
