package two_crystal_ball_problem

import "testing"

func TestTwoCrystalBalls(t *testing.T) {
	tests := []struct {
		name     string
		breaks   []bool
		expected int
	}{
		{
			name:     "Empty array",
			breaks:   []bool{},
			expected: -1,
		},
		{
			name:     "No breaks",
			breaks:   []bool{false, false, false, false, false},
			expected: -1,
		},
		{
			name:     "All breaks",
			breaks:   []bool{true, true, true, true, true},
			expected: 0,
		},
		{
			name:     "Single element - no break",
			breaks:   []bool{false},
			expected: -1,
		},
		{
			name:     "Single element - breaks",
			breaks:   []bool{true},
			expected: 0,
		},
		{
			name:     "Breaks at first floor",
			breaks:   []bool{true, true, true, true, true},
			expected: 0,
		},
		{
			name:     "Breaks at last floor",
			breaks:   []bool{false, false, false, false, true},
			expected: 4,
		},
		{
			name:     "Breaks in middle",
			breaks:   []bool{false, false, false, true, true, true},
			expected: 3,
		},
		{
			name:     "Larger array - breaks early",
			breaks:   []bool{false, false, true, true, true, true, true, true, true, true},
			expected: 2,
		},
		{
			name:     "Larger array - breaks late",
			breaks:   []bool{false, false, false, false, false, false, false, true, true, true},
			expected: 7,
		},
		{
			name:     "100 floors - breaks at 50",
			breaks:   make100FloorsBreakAt(50),
			expected: 50,
		},
		{
			name:     "100 floors - breaks at 1",
			breaks:   make100FloorsBreakAt(1),
			expected: 1,
		},
		{
			name:     "100 floors - breaks at 99",
			breaks:   make100FloorsBreakAt(99),
			expected: 99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoCrystalBalls(tt.breaks)
			if result != tt.expected {
				t.Errorf("TwoCrystalBalls() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func make100FloorsBreakAt(breakFloor int) []bool {
	floors := make([]bool, 100)
	for i := breakFloor; i < 100; i++ {
		floors[i] = true
	}
	return floors
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Error("Expected result to be a map")
	}

	if _, exists := resultMap["breaking_floor"]; !exists {
		t.Error("Expected result to contain 'breaking_floor' key")
	}
}

func BenchmarkTwoCrystalBalls10(b *testing.B) {
	breaks := make100FloorsBreakAt(5)[:10]

	for b.Loop() {
		TwoCrystalBalls(breaks)
	}
}

func BenchmarkTwoCrystalBalls100(b *testing.B) {
	breaks := make100FloorsBreakAt(50)

	for b.Loop() {
		TwoCrystalBalls(breaks)
	}
}

func BenchmarkTwoCrystalBalls1000(b *testing.B) {
	breaks := make([]bool, 1000)
	for i := 500; i < 1000; i++ {
		breaks[i] = true
	}

	for b.Loop() {
		TwoCrystalBalls(breaks)
	}
}

func BenchmarkTwoCrystalBalls10000(b *testing.B) {
	breaks := make([]bool, 10000)
	for i := 5000; i < 10000; i++ {
		breaks[i] = true
	}

	for b.Loop() {
		TwoCrystalBalls(breaks)
	}
}
