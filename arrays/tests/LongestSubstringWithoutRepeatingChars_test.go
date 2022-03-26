package arrays

import (
	"testing"

	arrays "github.com/ersindevrim/leetcode/arrays/code"
	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("This should return correct value 3", func(t *testing.T) {
		got := arrays.LengthOfLongestSubstring("abcabcbb")
		want := 3

		assert.Equal(t, got, want)
	})

	t.Run("This should return correct value 3", func(t *testing.T) {
		got := arrays.LengthOfLongestSubstring("bbbbb")
		want := 1

		assert.Equal(t, got, want)
	})

	t.Run("This should return correct value 1", func(t *testing.T) {
		got := arrays.LengthOfLongestSubstring(" ")
		want := 1

		assert.Equal(t, got, want)
	})

	t.Run("This should return correct value 0", func(t *testing.T) {
		got := arrays.LengthOfLongestSubstring("")
		want := 0

		assert.Equal(t, got, want)
	})

	t.Run("This should return correct value 3", func(t *testing.T) {
		got := arrays.LengthOfLongestSubstring("dvdf")
		want := 3

		assert.Equal(t, got, want)
	})
}
