package arrays

// DuplicateZeros ...
func DuplicateZeros(arr []int) []int {

	returnArray := make([]int, len(arr))
	howManyZeroesAdded := 0

	for i := 0; i < len(returnArray)-1; i++ {
		returnArray[i+howManyZeroesAdded] = arr[i]
		if arr[i] == 0 {
			howManyZeroesAdded++
			returnArray[i+howManyZeroesAdded] = 0
		}
		if i+howManyZeroesAdded == len(returnArray)-1 {
			break
		}
	}

	return returnArray
}
