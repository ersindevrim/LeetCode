package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("This should return fl", func(t *testing.T) {
		got := algorithm1.LongestCommonPrefix([]string{"flower", "flow", "flight"})
		want := "fl"
		assert.Equal(t, got, want)
	})

	t.Run("This should return empty string", func(t *testing.T) {
		got := algorithm1.LongestCommonPrefix([]string{"dog", "racecar", "car"})
		want := ""
		assert.Equal(t, got, want)
	})
}
