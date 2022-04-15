package problems

import (
	"testing"

	problems "github.com/ersindevrim/leetcode/problems/code"
	"github.com/stretchr/testify/assert"
)

type PalindromeCheck struct {
	Got  int
	Want bool
}

func TestIsPalindrome(t *testing.T) {
	t.Run("This should success ", func(t *testing.T) {
		qs := []PalindromeCheck{
			{
				Got:  123,
				Want: false,
			},
			{
				Got:  12321,
				Want: true,
			},
			{
				Got:  -121,
				Want: false,
			},
		}

		for _, v := range qs {
			assert.Equal(t, v.Want, problems.IsPalindrome(v.Got))
		}
	})
}
