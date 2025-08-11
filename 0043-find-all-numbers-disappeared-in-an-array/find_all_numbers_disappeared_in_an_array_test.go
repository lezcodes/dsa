package find_all_numbers_disappeared_in_an_array

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	nums     []int
	expected []int
}{
	{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	{[]int{1, 1}, []int{2}},
	{[]int{1, 2, 3, 4, 5}, []int{}},
	{[]int{2, 2, 2, 2}, []int{1, 3, 4}},
	{[]int{1}, []int{}},
	{[]int{2}, []int{1}},
	{[]int{1, 3, 3}, []int{2}},
}

func TestFindDisappearedNumbersHashSet(t *testing.T) {
	for _, test := range testCases {
		result := FindDisappearedNumbersHashSet(test.nums)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("FindDisappearedNumbersHashSet(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}

func TestFindDisappearedNumbersInPlace(t *testing.T) {
	for _, test := range testCases {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums)
		result := FindDisappearedNumbersInPlace(numsCopy)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("FindDisappearedNumbersInPlace(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}

func BenchmarkFindDisappearedNumbersHashSet(b *testing.B) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	for b.Loop() {
		FindDisappearedNumbersHashSet(nums)
	}
}

func BenchmarkFindDisappearedNumbersInPlace(b *testing.B) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	for b.Loop() {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		FindDisappearedNumbersInPlace(numsCopy)
	}
}
