package dungeon_game

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 1
	}
	n := len(dungeon[0])
	if n == 0 {
		return 1
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dp[i][j] = max(1, dp[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dp[i][j] = max(1, dp[i+1][j]-dungeon[i][j])
			} else {
				dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
			}
		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Run() any {
	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	return calculateMinimumHP(dungeon)
}
