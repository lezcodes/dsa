package max_subarray

type Result struct {
	MaxSum   int
	Subarray []int
	StartIdx int
	EndIdx   int
}

func maxSubArray(nums []int) Result {
	if len(nums) == 0 {
		return Result{MaxSum: 0, Subarray: []int{}, StartIdx: 0, EndIdx: 0}
	}

	maxSum := nums[0]
	currentSum := nums[0]
	start := 0
	end := 0
	tempStart := 0

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
			tempStart = i
		} else {
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
			start = tempStart
			end = i
		}
	}

	subarray := make([]int, end-start+1)
	copy(subarray, nums[start:end+1])

	return Result{
		MaxSum:   maxSum,
		Subarray: subarray,
		StartIdx: start,
		EndIdx:   end,
	}
}

func Run() any {
	testCases := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
		{-2, -3, -1, -5},
		{-1, -2, -3, -4},
		{1, 2, 3, 4, 5},
		{-5, -2, -8, -1},
	}

	results := make([]Result, len(testCases))
	for i, nums := range testCases {
		results[i] = maxSubArray(nums)
	}

	return results
}
