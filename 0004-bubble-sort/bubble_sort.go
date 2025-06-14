package bubble_sort

func BubbleSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	for i := range n - 1 {
		swapped := false
		for j := range n - i - 1 {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return result
}

func BubbleSortInPlace(arr []int) {
	n := len(arr)

	for i := range n - 1 {
		swapped := false
		for j := range n - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func BubbleSortStrings(arr []string) []string {
	n := len(arr)
	result := make([]string, n)
	copy(result, arr)

	for i := range n - 1 {
		swapped := false
		for j := range n - i - 1 {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return result
}

func BubbleSortFloat64(arr []float64) []float64 {
	n := len(arr)
	result := make([]float64, n)
	copy(result, arr)

	for i := range n - 1 {
		swapped := false
		for j := range n - i - 1 {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return result
}

func BubbleSortWithSteps(arr []int) ([]int, [][]int) {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	var steps [][]int
	steps = append(steps, make([]int, n))
	copy(steps[0], result)

	for i := range n - 1 {
		swapped := false
		for j := range n - i - 1 {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true

				step := make([]int, n)
				copy(step, result)
				steps = append(steps, step)
			}
		}
		if !swapped {
			break
		}
	}

	return result, steps
}

func Run() any {
	originalArray := []int{64, 34, 25, 12, 22, 11, 90}
	sortedArray := BubbleSort(originalArray)

	stringArray := []string{"banana", "apple", "cherry", "date"}
	sortedStrings := BubbleSortStrings(stringArray)

	floatArray := []float64{3.14, 2.71, 1.41, 0.57}
	sortedFloats := BubbleSortFloat64(floatArray)

	return map[string]any{
		"original_array":   originalArray,
		"sorted_array":     sortedArray,
		"original_strings": stringArray,
		"sorted_strings":   sortedStrings,
		"original_floats":  floatArray,
		"sorted_floats":    sortedFloats,
		"algorithm":        "Bubble Sort",
		"time_complexity":  "O(n²) average/worst, O(n) best",
		"space_complexity": "O(1)",
	}
}
