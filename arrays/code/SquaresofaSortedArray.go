package arrays

import "sort"

// SortedSquares ...
func SortedSquares(nums []int) []int {
	var returnArray = make([]int, len(nums))

	for i, v := range nums {
		returnArray[i] = v * v
	}

	sort.Slice(returnArray, func(i, j int) bool {
		return returnArray[i] < returnArray[j]
	})
	return returnArray
}
