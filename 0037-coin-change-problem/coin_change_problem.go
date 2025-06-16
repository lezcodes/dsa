package coin_change_problem

import "math"

type Result struct {
	Coins  []int
	Amount int
	Result int
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for _, coin := range coins {
		for x := coin; x <= amount; x++ {
			if dp[x-coin] != math.MaxInt32 {
				dp[x] = min(dp[x], dp[x-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Run() any {
	testCases := []struct {
		coins  []int
		amount int
	}{
		{[]int{1, 3, 4}, 6},
		{[]int{2}, 3},
		{[]int{1}, 0},
		{[]int{1, 2, 5}, 11},
		{[]int{2, 5, 10, 1}, 27},
		{[]int{5, 10, 25}, 30},
		{[]int{9, 6, 5, 1}, 11},
		{[]int{1, 5, 10, 21, 25}, 63},
	}

	results := make([]Result, len(testCases))
	for i, tc := range testCases {
		results[i] = Result{
			Coins:  tc.coins,
			Amount: tc.amount,
			Result: coinChange(tc.coins, tc.amount),
		}
	}

	return results
}
