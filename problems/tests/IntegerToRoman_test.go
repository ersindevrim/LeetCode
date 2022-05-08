package problems

import (
	"testing"

	problems "github.com/ersindevrim/leetcode/problems/code"
	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	t.Run("This should success ", func(t *testing.T) {
		got := problems.IntToRoman(1712)
		want := "MDCCXII"

		assert.Equal(t, got, want)

	})
	t.Run("This should success ", func(t *testing.T) {
		got := problems.IntToRoman(712)
		want := "DCCXII"

		assert.Equal(t, got, want)

	})
	t.Run("This should success ", func(t *testing.T) {
		got := problems.IntToRoman(69)
		want := "LXIX"

		assert.Equal(t, got, want)

	})
	t.Run("This should success ", func(t *testing.T) {
		got := problems.IntToRoman(12766)
		want := "MMMMMMMMMMMMDCCLXVI"

		assert.Equal(t, got, want)

	})
	t.Run("This should success ", func(t *testing.T) {
		got := problems.IntToRoman(1994)
		want := "MCMXCIV"

		assert.Equal(t, got, want)

	})
}
