package dungeon_game

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	testCases := []struct {
		dungeon [][]int
		want    int
	}{
		{
			dungeon: [][]int{
				{-2, -3, 3},
				{-5, -10, 1},
				{10, 30, -5},
			},
			want: 7,
		},
		{
			dungeon: [][]int{
				{0},
			},
			want: 1,
		},
		{
			dungeon: [][]int{
				{100},
			},
			want: 1,
		},
		{
			dungeon: [][]int{
				{-100},
			},
			want: 101,
		},
		{
			dungeon: [][]int{
				{0, 0, 0},
				{1, 1, -1},
			},
			want: 1,
		},
	}

	for _, tc := range testCases {
		got := calculateMinimumHP(tc.dungeon)
		if got != tc.want {
			t.Errorf("calculateMinimumHP(%v) = %d; want %d", tc.dungeon, got, tc.want)
		}
	}
}

func BenchmarkCalculateMinimumHP(b *testing.B) {
	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	for b.Loop() {
		calculateMinimumHP(dungeon)
	}
}
