package algorithm1

/* IsMatch ...
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​

'*' Matches zero or more of the preceding element

The matching should cover the entire input string (not partial).

Input: s = "aa", p = "a"

Output: false

Input: s = "aa", p = "a*"

Output: true

Input: s = "ab", p = ".*"

Output: true
*/
func IsMatch(myString string, pattern string) bool {
	switch {
	case len(pattern) == 0:
		return len(myString) == 0
	case len(myString) == 0:
		if len(pattern) > 1 && pattern[1] == '*' {
			return IsMatch(myString, pattern[2:])
		}
		return false
	case len(pattern) > 1 && pattern[1] == '*':
		if IsMatch(myString, pattern[2:]) {
			return true
		} else if myString[0] == pattern[0] || pattern[0] == '.' {
			return IsMatch(myString[1:], pattern)
		}
		return false
	default:
		return (myString[0] == pattern[0] || pattern[0] == '.') && IsMatch(myString[1:], pattern[1:])
	}
}
