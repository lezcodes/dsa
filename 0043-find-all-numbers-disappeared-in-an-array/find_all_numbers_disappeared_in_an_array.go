package find_all_numbers_disappeared_in_an_array

func FindDisappearedNumbersHashSet(nums []int) []int {
	seen := make(map[int]bool)
	for _, num := range nums {
		seen[num] = true
	}

	result := []int{}
	for i := 1; i <= len(nums); i++ {
		if !seen[i] {
			result = append(result, i)
		}
	}

	return result
}

func FindDisappearedNumbersInPlace(nums []int) []int {
	for _, num := range nums {
		index := abs(num) - 1
		if index >= 0 && index < len(nums) && nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	result := []int{}
	for i := range nums {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Run() any {
	examples := [][]int{
		{4, 3, 2, 7, 8, 2, 3, 1},
		{1, 1},
		{1, 2, 3, 4, 5},
		{2, 2, 2, 2},
	}

	hashResults := make([][]int, len(examples))
	inPlaceResults := make([][]int, len(examples))

	for i, nums := range examples {
		hashResults[i] = FindDisappearedNumbersHashSet(nums)

		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		inPlaceResults[i] = FindDisappearedNumbersInPlace(numsCopy)
	}

	return map[string]any{
		"hashSet": hashResults,
		"inPlace": inPlaceResults,
	}
}
