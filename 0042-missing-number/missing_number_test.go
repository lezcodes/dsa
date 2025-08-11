package missing_number

import "testing"

var testCases = []struct {
	nums     []int
	expected int
}{
	{[]int{3, 0, 1}, 2},
	{[]int{0, 1}, 2},
	{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	{[]int{1}, 0},
	{[]int{0}, 1},
	{[]int{1, 2}, 0},
	{[]int{0, 2}, 1},
}

func TestMissingNumberSum(t *testing.T) {
	for _, test := range testCases {
		result := MissingNumberSum(test.nums)
		if result != test.expected {
			t.Errorf("MissingNumberSum(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}

func TestMissingNumberXOR(t *testing.T) {
	for _, test := range testCases {
		result := MissingNumberXOR(test.nums)
		if result != test.expected {
			t.Errorf("MissingNumberXOR(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}

func BenchmarkMissingNumberSum(b *testing.B) {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	for b.Loop() {
		MissingNumberSum(nums)
	}
}

func BenchmarkMissingNumberXOR(b *testing.B) {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	for b.Loop() {
		MissingNumberXOR(nums)
	}
}
