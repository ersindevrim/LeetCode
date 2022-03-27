package strings

import (
	"testing"

	strings "github.com/ersindevrim/leetcode/strings/code"
	"github.com/stretchr/testify/assert"
)

type question struct {
	got  string
	want string
}

func TestMoveZeroes(t *testing.T) {
	t.Run("This should success ", func(t *testing.T) {
		qs := []question{
			{
				got:  "babad",
				want: "bab",
			},
			{
				got:  "cbbd",
				want: "bb",
			},
			{
				got:  "abbcccddcccbba",
				want: "abbcccddcccbba",
			},
			{
				got:  "a",
				want: "a",
			},
		}

		for _, v := range qs {
			assert.Equal(t, strings.LongestPalindrome(v.got), v.want)
		}
	})
}
