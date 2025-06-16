package coin_change_problem

import "testing"

func TestCoinChange(t *testing.T) {
	testCases := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 3, 4}, 6, 2},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2, 5, 10, 1}, 27, 4},
		{[]int{5, 10, 25}, 30, 2},
		{[]int{9, 6, 5, 1}, 11, 2},
		{[]int{1, 5, 10, 21, 25}, 63, 3},
	}

	for i, tc := range testCases {
		result := coinChange(tc.coins, tc.amount)
		if result != tc.expected {
			t.Errorf("Test case %d: expected %d coins, got %d", i+1, tc.expected, result)
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

	expectedResults := []int{2, -1, 0, 3, 4, 2, 2, 3}
	for i, res := range results {
		if i < len(expectedResults) && res.Result != expectedResults[i] {
			t.Errorf("Result %d: expected %d, got %d", i, expectedResults[i], res.Result)
		}
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}

func BenchmarkCoinChange(b *testing.B) {
	coins := []int{1, 2, 5}
	amount := 11

	for b.Loop() {
		coinChange(coins, amount)
	}
}
