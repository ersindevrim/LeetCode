package problems

// Reverse ...
func Reverse(x int) int {
	answer := 0
	shouldRun := x

	if x > int(2147483647) || x < -int(2147483647) {
		return 0
	}

	if x < 0 {
		shouldRun = -shouldRun
	}
	for shouldRun > 0 {
		remainder := shouldRun % 10
		answer *= 10
		answer += remainder
		shouldRun /= 10
	}

	if answer > int(2147483647) {
		return 0
	}

	if x < 0 {
		return -answer
	}
	return answer
}
