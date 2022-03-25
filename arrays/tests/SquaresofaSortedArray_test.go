package arrays

import (
	"testing"

	arrays "github.com/ersindevrim/leetcode/arrays/code"
	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	t.Run("This should return correct value 2", func(t *testing.T) {
		got := arrays.SortedSquares([]int{-4, -1, 0, 3, 10})
		want := []int{0, 1, 9, 16, 100}

		assert.Equal(t, got, want)
	})

}
