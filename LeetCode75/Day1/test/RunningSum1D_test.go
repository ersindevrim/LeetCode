package LeetCode75

import (
	"testing"

	leetCode "github.com/ersindevrim/leetcode/LeetCode75/Day1/Code"
	"github.com/stretchr/testify/assert"
)

func TestRunningSum(t *testing.T) {
	t.Run("This should return -1 for non found search", func(t *testing.T) {

		got := leetCode.RunningSum([]int{1, 2, 3, 4})
		want := []int{1, 3, 6, 10}

		assert.Equal(t, want, got)
	})

}
