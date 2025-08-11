package missing_number

func MissingNumberSum(nums []int) int {
	n := len(nums)
	expectedSum := n * (n + 1) / 2
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return expectedSum - actualSum
}

func MissingNumberXOR(nums []int) int {
	n := len(nums)
	result := n
	for i, num := range nums {
		result ^= i ^ num
	}
	return result
}

func Run() any {
	examples := [][]int{
		{3, 0, 1},
		{0, 1},
		{9, 6, 4, 2, 3, 5, 7, 0, 1},
	}

	sumResults := make([]int, len(examples))
	xorResults := make([]int, len(examples))

	for i, nums := range examples {
		sumResults[i] = MissingNumberSum(nums)
		xorResults[i] = MissingNumberXOR(nums)
	}

	return map[string]any{
		"sum": sumResults,
		"xor": xorResults,
	}
}
