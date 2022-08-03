package problems

import (
	"testing"

	problems "github.com/ersindevrim/leetcode/problems/code"
	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	t.Run("This should success ", func(t *testing.T) {
		got := problems.RomanToInt("MDCCXII")
		want := 1712

		assert.Equal(t, got, want)

	})
	t.Run("This should success ", func(t *testing.T) {
		got := problems.RomanToInt("DCCXII")
		want := 712

		assert.Equal(t, got, want)

	})
	t.Run("This should success ", func(t *testing.T) {
		got := problems.RomanToInt("LXIX")
		want := 69

		assert.Equal(t, got, want)

	})
	t.Run("This should success ", func(t *testing.T) {
		got := problems.RomanToInt("MMMMMMMMMMMMDCCLXVI")
		want := 12766

		assert.Equal(t, got, want)

	})
	t.Run("This should success ", func(t *testing.T) {
		got := problems.RomanToInt("MCMXCIV")
		want := 1994

		assert.Equal(t, got, want)

	})
}
