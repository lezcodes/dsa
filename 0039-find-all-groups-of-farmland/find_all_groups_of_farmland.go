package find_all_groups_of_farmland

type FarmlandGroup struct {
	TopLeft     [2]int
	BottomRight [2]int
	Area        int
}

type TestCase struct {
	Land   [][]int
	Groups [][]int
}

func findFarmland(land [][]int) [][]int {
	if len(land) == 0 || len(land[0]) == 0 {
		return [][]int{}
	}

	m, n := len(land), len(land[0])
	result := [][]int{}

	landCopy := make([][]int, m)
	for i := range land {
		landCopy[i] = make([]int, n)
		copy(landCopy[i], land[i])
	}

	for row1 := range m {
		for col1 := range n {
			if landCopy[row1][col1] == 1 {
				x, y := row1, col1

				for x < m && landCopy[x][col1] == 1 {
					y = col1
					for y < n && landCopy[x][y] == 1 {
						landCopy[x][y] = 0
						y++
					}
					x++
				}

				result = append(result, []int{row1, col1, x - 1, y - 1})
			}
		}
	}

	return result
}

func findFarmlandWithDetails(land [][]int) []FarmlandGroup {
	if len(land) == 0 || len(land[0]) == 0 {
		return []FarmlandGroup{}
	}

	m, n := len(land), len(land[0])
	result := []FarmlandGroup{}

	landCopy := make([][]int, m)
	for i := range land {
		landCopy[i] = make([]int, n)
		copy(landCopy[i], land[i])
	}

	for row1 := range m {
		for col1 := range n {
			if landCopy[row1][col1] == 1 {
				x, y := row1, col1

				for x < m && landCopy[x][col1] == 1 {
					y = col1
					for y < n && landCopy[x][y] == 1 {
						landCopy[x][y] = 0
						y++
					}
					x++
				}

				group := FarmlandGroup{
					TopLeft:     [2]int{row1, col1},
					BottomRight: [2]int{x - 1, y - 1},
					Area:        (x - row1) * (y - col1),
				}
				result = append(result, group)
			}
		}
	}

	return result
}

func Run() any {
	testCases := []TestCase{
		{
			Land: [][]int{
				{1, 0, 0},
				{0, 1, 1},
				{0, 1, 1},
			},
			Groups: [][]int{{0, 0, 0, 0}, {1, 1, 2, 2}},
		},
		{
			Land: [][]int{
				{1, 1},
				{1, 1},
			},
			Groups: [][]int{{0, 0, 1, 1}},
		},
		{
			Land: [][]int{
				{0},
			},
			Groups: [][]int{},
		},
		{
			Land: [][]int{
				{1, 1, 1, 1, 0, 0},
				{1, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 1, 1},
				{0, 0, 0, 0, 1, 1},
			},
			Groups: [][]int{{0, 0, 1, 3}, {2, 4, 3, 5}},
		},
		{
			Land: [][]int{
				{1, 0, 1},
				{0, 1, 0},
				{1, 0, 1},
			},
			Groups: [][]int{{0, 0, 0, 0}, {0, 2, 0, 2}, {1, 1, 1, 1}, {2, 0, 2, 0}, {2, 2, 2, 2}},
		},
	}

	results := make([][][]int, len(testCases))
	for i, tc := range testCases {
		results[i] = findFarmland(tc.Land)
	}

	return results
}
