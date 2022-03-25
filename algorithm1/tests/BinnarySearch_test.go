package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/stretchr/testify/assert"
)

func TestBinnarySearch(t *testing.T) {
	t.Run("This should return -1 for non found search", func(t *testing.T) {
		got := algorithm1.Search([]int{2, 7, 11, 15}, 9)
		want := -1

		assert.Equal(t, got, want)
	})

	t.Run("This should return index for found search", func(t *testing.T) {
		got := algorithm1.Search([]int{2, 7, 11, 15}, 7)
		want := 1

		assert.Equal(t, got, want)
	})
}
