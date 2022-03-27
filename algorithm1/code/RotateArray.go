package algorithm1

// Rotate ...
func Rotate(nums []int, k int) []int {
	startArray := nums[len(nums)-k:]
	endArray := nums[:len(nums)-k]

	returnArry := append(startArray, endArray...)

	return returnArry
}
