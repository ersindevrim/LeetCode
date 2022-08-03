package problems

// RomanIntMap ...
var RomanIntMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

/* RomanToInt ...
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func RomanToInt(s string) int {
	if s == "" {
		return 0
	}
	num, lastInt, total := 0, 0, 0

	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = RomanIntMap[char]
		if num < lastInt {
			total = total - num
		} else {
			total = total + num
		}
		lastInt = num
	}
	return total
}
