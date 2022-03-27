package algorithm1

// TwoSum ...
func TwoSum(numbers []int, target int) []int {

	for i, v := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if v+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{0}
}
