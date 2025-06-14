package linear_search

func Run() any {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	target := 22

	result := SearchInt(arr, target)
	if result != -1 {
		return map[string]any{
			"array":  arr,
			"target": target,
			"index":  result,
			"found":  true,
		}
	}

	return map[string]any{
		"array":  arr,
		"target": target,
		"index":  -1,
		"found":  false,
	}
}

func SearchInt(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

func SearchString(arr []string, target string) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

func SearchFloat64(arr []float64, target float64) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

func SearchGeneric[T comparable](arr []T, target T) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}
