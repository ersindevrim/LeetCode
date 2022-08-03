package arrays

// LengthOfLongestSubstring ...
func LengthOfLongestSubstring(s string) int {
	var i, maximumLen int = 0, 0
	mymap := make(map[byte]int)

	for j := 0; j < len(s); j++ {
		mymap[s[j]]++
		for mymap[s[j]] == 2 && i < j {
			mymap[s[i]]--
			i++
		}

		if maximumLen < j-i+1 {
			maximumLen = j - i + 1
		}
	}

	return maximumLen
}
