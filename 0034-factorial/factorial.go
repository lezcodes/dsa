package factorial

func Run() any {
	n := 5
	iterative := FactorialIterative(n)
	recursive := FactorialRecursive(n)

	return map[string]any{
		"n":         n,
		"iterative": iterative,
		"recursive": recursive,
	}
}

func FactorialIterative(n int) int {
	if n <= 1 {
		return 1
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func FactorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}
