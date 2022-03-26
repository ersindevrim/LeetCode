package arrays

import (
	"testing"

	arrays "github.com/ersindevrim/leetcode/arrays/code"
	"github.com/stretchr/testify/assert"
)

func TestDuplicateZeros(t *testing.T) {
	t.Run("This should return correct value 2", func(t *testing.T) {
		got := arrays.DuplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
		want := []int{1, 0, 0, 2, 3, 0, 0, 4}

		assert.Equal(t, got, want)
	})
}
