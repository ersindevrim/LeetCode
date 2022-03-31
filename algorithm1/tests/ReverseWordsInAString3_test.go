package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	t.Run("This should return correct", func(t *testing.T) {
		got := algorithm1.ReverseWords("a b d ee$#% aef$ea eaef eaj ae##ea'fe")
		want := "a b d %#$ee ae$fea feae jae ef'ae##ea"

		assert.Equal(t, got, want)
	})

	t.Run("This should return correct", func(t *testing.T) {
		got := algorithm1.ReverseWords("Let's take LeetCode contest")
		want := "s'teL ekat edoCteeL tsetnoc"

		assert.Equal(t, got, want)
	})

	t.Run("This should return correct", func(t *testing.T) {
		got := algorithm1.ReverseWords("git gud")
		want := "tig dug"

		assert.Equal(t, got, want)
	})
}
