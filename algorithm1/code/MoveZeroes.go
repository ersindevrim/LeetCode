package algorithm1

// MoveZeroes ...
func MoveZeroes(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}

	return nums
}
