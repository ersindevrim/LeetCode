package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	t.Run("This should return 1,3,12,0,0", func(t *testing.T) {
		got := algorithm1.MoveZeroes([]int{0, 1, 0, 3, 12})
		want := []int{1, 3, 12, 0, 0}

		assert.Equal(t, got, want)
	})

	t.Run("This should return 0", func(t *testing.T) {
		got := algorithm1.MoveZeroes([]int{0})
		want := []int{0}

		assert.Equal(t, got, want)
	})

}
