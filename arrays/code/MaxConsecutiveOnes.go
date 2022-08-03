package arrays

// FindMaxConsecutiveOnes ...
func FindMaxConsecutiveOnes(nums []int) int {

	consecutive := 0
	find := 0

	for _, v := range nums {
		if v == 1 {
			find++
		} else {
			if find > consecutive {
				consecutive = find
			}
			find = 0
		}
	}

	if find > consecutive {
		consecutive = find
	}
	return consecutive
}
