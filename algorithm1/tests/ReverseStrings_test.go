package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	t.Run("This should return correct", func(t *testing.T) {
		got := algorithm1.ReverseString([]byte{'h', 'e', 'l', 'l', 'o'})
		want := []byte{'o', 'l', 'l', 'e', 'h'}

		assert.Equal(t, got, want)
	})

}
