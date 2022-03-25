package arrays

import (
	"testing"

	arrays "github.com/ersindevrim/leetcode/arrays/code"
	"github.com/stretchr/testify/assert"
)

func TestBinnarySearch(t *testing.T) {
	t.Run("This should return correct value 3", func(t *testing.T) {
		got := arrays.FindMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
		want := 3

		assert.Equal(t, got, want)
	})

	t.Run("This should return correct value 6", func(t *testing.T) {
		got := arrays.FindMaxConsecutiveOnes([]int{1, 1, 1, 1, 1, 1, 0, 1, 1, 1})
		want := 3

		assert.Equal(t, got, want)
	})
}
