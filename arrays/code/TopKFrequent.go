package arrays

import "sort"

type mySlice struct {
	Key   int
	Value int
}

// TopKFrequent ...
func TopKFrequent(nums []int, k int) []int {

	freq := make(map[int]int)
	for _, num := range nums {
		freq[num] += 1
	}

	var slice []mySlice
	for k, v := range freq {
		slice = append(slice, mySlice{k, v})
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Value > slice[j].Value
	})

	index := 1
	keys := make([]int, 0, k)
	for _, value := range slice {
		keys = append(keys, value.Key)
		if index == k {
			break
		}
		index++
	}

	return keys
}
