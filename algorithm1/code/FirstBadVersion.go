package algorithm1

func FirstBadVersion(n int) int {
	start := 0
	stop := n
	middle := 0

	position := 1

	for start <= stop {
		middle = start + (stop-start)/2

		//This is an api call from LEETCODE. To simulate that i will take bad version from function.
		isBad := isBadVersion(middle)
		if isBad {
			position = middle
			stop = middle - 1
		} else {
			start = middle + 1
		}
	}

	return position
}

func isBadVersion(version int) bool {
	return version == 3
}
