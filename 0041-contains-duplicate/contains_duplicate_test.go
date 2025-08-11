package contains_duplicate

import "testing"

var testCases = []struct {
	nums     []int
	expected bool
}{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	{[]int{}, false},
	{[]int{1}, false},
	{[]int{1, 1}, true},
}

func TestContainsDuplicate(t *testing.T) {
	for _, test := range testCases {
		result := ContainsDuplicate(test.nums)
		if result != test.expected {
			t.Errorf("ContainsDuplicate(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}

func TestContainsDuplicateInPlace(t *testing.T) {
	for _, test := range testCases {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums)
		result := ContainsDuplicateInPlace(numsCopy)
		if result != test.expected {
			t.Errorf("ContainsDuplicateInPlace(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	for b.Loop() {
		ContainsDuplicate(nums)
	}
}

func BenchmarkContainsDuplicateInPlace(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	for b.Loop() {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		ContainsDuplicateInPlace(numsCopy)
	}
}
