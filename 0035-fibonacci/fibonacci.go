package fibonacci

func Run() any {
	n := 10
	iterative := FibonacciIterative(n)
	recursive := FibonacciRecursive(n)

	return map[string]any{
		"n":         n,
		"iterative": iterative,
		"recursive": recursive,
	}
}

func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	prev, current := 0, 1
	for i := 2; i <= n; i++ {
		prev, current = current, prev+current
	}

	return current
}

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}
