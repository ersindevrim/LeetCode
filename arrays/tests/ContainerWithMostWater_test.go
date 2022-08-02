package arrays

import (
	"testing"

	arrays "github.com/ersindevrim/leetcode/arrays/code"
	"github.com/stretchr/testify/assert"
)

func TestContainerWithMostWater(t *testing.T) {
	t.Run("This should return correct value 16", func(t *testing.T) {
		got := arrays.MaxArea([]int{4, 3, 2, 1, 4})
		want := 16

		assert.Equal(t, got, want)
	})

	t.Run("This should return correct value 8", func(t *testing.T) {
		got := arrays.MaxArea([]int{1, 0, 0, 0, 0, 0, 0, 2, 2})
		want := 8

		assert.Equal(t, got, want)
	})

	t.Run("This should return correct value 0", func(t *testing.T) {
		got := arrays.MaxArea([]int{2, 0})
		want := 0

		assert.Equal(t, got, want)
	})

	t.Run("This should return correct value 80", func(t *testing.T) {
		got := arrays.MaxArea([]int{8, 10, 14, 0, 13, 10, 9, 9, 11, 11})
		want := 80

		assert.Equal(t, got, want)
	})
}
