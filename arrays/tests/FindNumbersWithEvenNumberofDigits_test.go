package arrays

import (
	"testing"

	arrays "github.com/ersindevrim/leetcode/arrays/code"
	"github.com/stretchr/testify/assert"
)

func TestFindNumbers(t *testing.T) {
	t.Run("This should return correct value 2", func(t *testing.T) {
		got := arrays.FindNumbers([]int{12, 345, 2, 6, 7896})
		want := 2

		assert.Equal(t, got, want)
	})

}
