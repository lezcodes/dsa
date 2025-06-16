package linear_search

func Run() any {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	target := 22
	index := Search(arr, target)
	return map[string]any{
		"array":  arr,
		"target": target,
		"index":  index,
	}
}

func Search[T comparable](arr []T, target T) int {
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
