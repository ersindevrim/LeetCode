package LeetCode75

// PivotIndex ...
/*
 * Input: nums = [1,7,3,6,5,6]
 * Output: 3
 * Explanation:
 * The pivot index is 3.
 * Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
 * Right sum = nums[4] + nums[5] = 5 + 6 = 11
 */
func PivotIndex(nums []int) int {
	for index, _ := range nums {
		var leftTotal, rightTotal int
		//To Left
		for i := index - 1; i >= 0; i-- {
			leftTotal += nums[i]
		}
		//To Right
		for i := index + 1; i <= len(nums)-1; i++ {
			rightTotal += nums[i]
		}

		if leftTotal == rightTotal {
			return index
		}
	}
	return -1
}
