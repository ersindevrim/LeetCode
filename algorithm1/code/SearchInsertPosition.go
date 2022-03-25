package algorithm1

func SearchInsert(nums []int, target int) int {
	insertPos := 0
	lenOfNumbers := len(nums)

	if target > nums[lenOfNumbers-1] {
		return lenOfNumbers
	}

	for i, v := range nums {
		if v == target {
			return i
		}

		//so we don't have target in given array
		if nums[i] > target {
			insertPos = i
			break
		}
	}

	return insertPos
}
