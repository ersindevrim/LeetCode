package LeetCode75

// RunningSum ...
func RunningSum(nums []int) []int {
	result := make([]int, len(nums))

	for index := range nums {
		for i := 0; i <= index; i++ {
			result[index] = result[index] + nums[i]
		}

	}

	return result
}
