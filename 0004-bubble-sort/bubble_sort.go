package bubble_sort

func Run() any {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	sorted := BubbleSortInt(arr)
	return map[string]any{
		"original": arr,
		"sorted":   sorted,
	}
}

func BubbleSort[T comparable](arr []T, less func(T, T) bool) []T {
	n := len(arr)
	result := make([]T, n)
	copy(result, arr)

	for i := range n - 1 {
		swapped := false
		for j := range n - i - 1 {
			if !less(result[j], result[j+1]) && result[j] != result[j+1] {
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

func BubbleSortInt(arr []int) []int {
	return BubbleSort(arr, func(a, b int) bool { return a < b })
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
