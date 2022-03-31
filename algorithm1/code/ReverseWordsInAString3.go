package algorithm1

import (
	"strings"
)

// ReverseWords ...
func ReverseWords(s string) string {
	strArr := strings.Split(s, " ")

	for k, v := range strArr {
		s := []byte(v)

		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}

		strArr[k] = string(s)
	}
	return strings.Join(strArr, " ")
}
