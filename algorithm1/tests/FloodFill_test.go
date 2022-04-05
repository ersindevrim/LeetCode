package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/stretchr/testify/assert"
)

func TestFloodFill(t *testing.T) {
	t.Run("This should return correct version", func(t *testing.T) {
		got := algorithm1.FloodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 2)
		want := [][]int{{2, 2, 2}, {2, 2, 2}}

		assert.Equal(t, got, want)
	})
}
