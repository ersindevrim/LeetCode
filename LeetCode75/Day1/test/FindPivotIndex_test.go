package LeetCode75

import (
	"testing"

	leetCode "github.com/ersindevrim/leetcode/LeetCode75/Day1/Code"
	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	t.Run("This should return -1 for non found search", func(t *testing.T) {

		got := leetCode.PivotIndex([]int{1, 7, 3, 6, 5, 6})
		want := 3

		assert.Equal(t, want, got)
	})

}
