package algorithm1

// Search ...
func Search(nums []int, target int) int {
	for index, value := range nums {
		if value == target {
			return index
		}
	}

	return -1
}
