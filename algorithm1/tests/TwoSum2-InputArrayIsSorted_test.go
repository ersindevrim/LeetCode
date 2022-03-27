package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Run("This should return 1,2", func(t *testing.T) {
		got := algorithm1.TwoSum([]int{2, 7, 11, 15}, 9)
		want := []int{1, 2}

		assert.Equal(t, got, want)
	})

	t.Run("This should return 1,2", func(t *testing.T) {
		got := algorithm1.TwoSum([]int{2, 3, 4}, 6)
		want := []int{1, 3}

		assert.Equal(t, got, want)
	})

	t.Run("This should return 1,2", func(t *testing.T) {
		got := algorithm1.TwoSum([]int{-1, 0}, -1)
		want := []int{1, 2}

		assert.Equal(t, got, want)
	})

}
