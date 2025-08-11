package contains_duplicate

func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		seen[num] = true
	}
	return len(seen) != len(nums)
}

func ContainsDuplicateInPlace(nums []int) bool {
	for i := range nums {
		index := abs(nums[i])
		if index > len(nums) {
			continue
		}
		if index > 0 && index <= len(nums) {
			if nums[index-1] < 0 {
				return true
			}
			nums[index-1] = -nums[index-1]
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Run() any {
	examples := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}

	hashResults := make([]bool, len(examples))
	for i, nums := range examples {
		hashResults[i] = ContainsDuplicate(nums)
	}

	inPlaceResults := make([]bool, len(examples))
	for i, nums := range examples {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		inPlaceResults[i] = ContainsDuplicateInPlace(numsCopy)
	}

	return map[string]any{
		"hashSet": hashResults,
		"inPlace": inPlaceResults,
	}
}
