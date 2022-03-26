package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	t.Run("This should return correct value 3", func(t *testing.T) {
		got := algorithm1.Rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
		want := []int{5, 6, 7, 1, 2, 3, 4}

		assert.Equal(t, got, want)
	})

}
