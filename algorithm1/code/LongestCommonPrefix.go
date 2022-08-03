package algorithm1

/* LongestCommonPrefix ...

Input: strs = ["flower","flow","flight"]
Output: "fl"

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/
func LongestCommonPrefix(strs []string) string {
	arrLen := len(strs)
	if len(strs) == 0 {
		return ""
	}

	minLen := len(strs[0])
	tmp := 0

	for i := 1; i < arrLen; i++ {
		tmp = len(strs[i])
		if tmp < minLen {
			minLen = len(strs[i])
		}
	}

	ans := ""
	for j := 0; j < minLen; j++ {
		for k := 0; k < arrLen-1; k++ {
			if strs[k][j] != strs[k+1][j] {
				return ans
			}
		}
		ans += string(strs[0][j])
	}
	return ans
}
