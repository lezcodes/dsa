package max_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-2, -3, -1, -5}, -1},
		{[]int{-1, -2, -3, -4}, -1},
		{[]int{1, 2, 3, 4, 5}, 15},
	}

	for i, tc := range testCases {
		result := maxSubArray(tc.input)
		if result.MaxSum != tc.expected {
			t.Errorf("Test case %d: expected max sum %d, got %d", i+1, tc.expected, result.MaxSum)
		}

		sum := 0
		for _, val := range result.Subarray {
			sum += val
		}
		if sum != result.MaxSum {
			t.Errorf("Test case %d: subarray sum %d doesn't match max sum %d", i+1, sum, result.MaxSum)
		}
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	results, ok := result.([]Result)
	if !ok {
		t.Error("Expected result to be []Result")
		return
	}

	if len(results) == 0 {
		t.Error("Expected at least one result")
	}

	for i, res := range results {
		sum := 0
		for _, val := range res.Subarray {
			sum += val
		}
		if sum != res.MaxSum {
			t.Errorf("Result %d: subarray sum %d doesn't match max sum %d", i, sum, res.MaxSum)
		}
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	for b.Loop() {
		maxSubArray(nums)
	}
}
