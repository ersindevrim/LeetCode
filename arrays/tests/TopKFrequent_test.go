package arrays

import (
	"testing"

	arrays "github.com/ersindevrim/leetcode/arrays/code"
	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	t.Run("This should return correct value 16", func(t *testing.T) {
		got := arrays.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
		want := []int{1, 2}

		assert.Equal(t, got, want)
	})
}
