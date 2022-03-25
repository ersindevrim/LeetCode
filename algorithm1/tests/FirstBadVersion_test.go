package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/stretchr/testify/assert"
)

func TestFirstBadVersion(t *testing.T) {
	t.Run("This should return correct version", func(t *testing.T) {
		got := algorithm1.FirstBadVersion(6)
		want := 3

		assert.Equal(t, got, want)
	})
}
