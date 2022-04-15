package problems

func IsPalindrome(x int) bool {
	answer := 0
	shouldRun := x

	if x < 0 {
		shouldRun = -shouldRun
	}
	for shouldRun > 0 {
		remainder := shouldRun % 10
		answer *= 10
		answer += remainder
		shouldRun /= 10
	}

	if x < 0 {
		return false
	}
	return answer == x
}
