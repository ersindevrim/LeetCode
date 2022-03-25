package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	t.Run("This should return inserted location if given number in middle of list", func(t *testing.T) {
		got := algorithm1.SearchInsert([]int{1, 3, 5, 6}, 2)
		want := 1

		assert.Equal(t, got, want)
	})

	t.Run("This should return location of given number", func(t *testing.T) {
		got := algorithm1.SearchInsert([]int{1, 3, 5, 6}, 5)
		want := 2

		assert.Equal(t, got, want)
	})

	t.Run("This should return inserted location if given number out of list", func(t *testing.T) {
		got := algorithm1.SearchInsert([]int{1, 3, 5, 6}, 7)
		want := 4

		assert.Equal(t, got, want)
	})
}
