package strings

// Convert ...
func Convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	i, row := 0, 0
	down := true
	array := make([]string, numRows)
	for i = 0; i < len(s); i++ {
		array[row] += string(s[i])
		if row == 0 {
			down = true
		}
		if row == numRows-1 {
			down = false
		}
		if down {
			row++
		} else {
			row--
		}
	}
	answer := ""
	for i = 0; i < numRows; i++ {
		answer += array[i]
	}
	return answer
}
