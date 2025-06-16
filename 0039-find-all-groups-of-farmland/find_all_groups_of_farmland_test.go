package find_all_groups_of_farmland

import (
	"reflect"
	"sort"
	"testing"
)

func sortGroups(groups [][]int) {
	sort.Slice(groups, func(i, j int) bool {
		if groups[i][0] != groups[j][0] {
			return groups[i][0] < groups[j][0]
		}
		if groups[i][1] != groups[j][1] {
			return groups[i][1] < groups[j][1]
		}
		if groups[i][2] != groups[j][2] {
			return groups[i][2] < groups[j][2]
		}
		return groups[i][3] < groups[j][3]
	})
}

func TestFindFarmland(t *testing.T) {
	testCases := []struct {
		name     string
		land     [][]int
		expected [][]int
	}{
		{
			name: "Single cell farmland",
			land: [][]int{
				{1, 0, 0},
				{0, 1, 1},
				{0, 1, 1},
			},
			expected: [][]int{{0, 0, 0, 0}, {1, 1, 2, 2}},
		},
		{
			name: "Full rectangle",
			land: [][]int{
				{1, 1},
				{1, 1},
			},
			expected: [][]int{{0, 0, 1, 1}},
		},
		{
			name: "No farmland",
			land: [][]int{
				{0},
			},
			expected: [][]int{},
		},
		{
			name: "Multiple separate rectangles",
			land: [][]int{
				{1, 1, 1, 1, 0, 0},
				{1, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 1, 1},
				{0, 0, 0, 0, 1, 1},
			},
			expected: [][]int{{0, 0, 1, 3}, {2, 4, 3, 5}},
		},
		{
			name: "Individual cells",
			land: [][]int{
				{1, 0, 1},
				{0, 1, 0},
				{1, 0, 1},
			},
			expected: [][]int{{0, 0, 0, 0}, {0, 2, 0, 2}, {1, 1, 1, 1}, {2, 0, 2, 0}, {2, 2, 2, 2}},
		},
		{
			name:     "Empty matrix",
			land:     [][]int{},
			expected: [][]int{},
		},
		{
			name:     "Single row",
			land:     [][]int{{1, 1, 0, 1}},
			expected: [][]int{{0, 0, 0, 1}, {0, 3, 0, 3}},
		},
		{
			name: "Single column",
			land: [][]int{
				{1},
				{1},
				{0},
				{1},
			},
			expected: [][]int{{0, 0, 1, 0}, {3, 0, 3, 0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findFarmland(tc.land)

			sortGroups(result)
			sortGroups(tc.expected)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestFindFarmlandWithDetails(t *testing.T) {
	land := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}

	result := findFarmlandWithDetails(land)

	if len(result) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(result))
	}

	for _, group := range result {
		if group.Area <= 0 {
			t.Errorf("Expected positive area, got %d", group.Area)
		}

		if group.TopLeft[0] > group.BottomRight[0] || group.TopLeft[1] > group.BottomRight[1] {
			t.Errorf("Invalid group bounds: top-left %v, bottom-right %v", group.TopLeft, group.BottomRight)
		}
	}
}

func TestGroupValidation(t *testing.T) {
	land := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	result := findFarmland(land)

	if len(result) != 1 {
		t.Errorf("Expected 1 group, got %d", len(result))
	}

	if len(result) > 0 {
		group := result[0]
		expectedWidth := group[3] - group[1] + 1
		expectedHeight := group[2] - group[0] + 1
		expectedArea := expectedWidth * expectedHeight

		if expectedArea != 9 {
			t.Errorf("Expected area 9, calculated %d", expectedArea)
		}
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	results, ok := result.([][][]int)
	if !ok {
		t.Error("Expected result to be [][][]int")
		return
	}

	if len(results) == 0 {
		t.Error("Expected at least one result")
	}

	for i, res := range results {
		for j, group := range res {
			if len(group) != 4 {
				t.Errorf("Result %d, group %d: expected 4 coordinates, got %d", i, j, len(group))
			}

			if len(group) == 4 {
				r1, c1, r2, c2 := group[0], group[1], group[2], group[3]
				if r1 > r2 || c1 > c2 {
					t.Errorf("Result %d, group %d: invalid bounds [%d,%d,%d,%d]", i, j, r1, c1, r2, c2)
				}
			}
		}
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}

func BenchmarkFindFarmland(b *testing.B) {
	land := [][]int{
		{1, 1, 1, 1, 0, 0, 1, 1},
		{1, 1, 1, 1, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0},
		{1, 0, 1, 0, 0, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 1, 1},
	}

	b.ResetTimer()
	for b.Loop() {
		findFarmland(land)
	}
}

func BenchmarkFindFarmlandWithDetails(b *testing.B) {
	land := [][]int{
		{1, 1, 1, 1, 0, 0, 1, 1},
		{1, 1, 1, 1, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0},
	}

	b.ResetTimer()
	for b.Loop() {
		findFarmlandWithDetails(land)
	}
}
