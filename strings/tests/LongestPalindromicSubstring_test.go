package strings

import (
	"testing"

	"github.com/ersindevrim/leetcode/global"
	strings "github.com/ersindevrim/leetcode/strings/code"
	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	t.Run("This should success ", func(t *testing.T) {

		qs := []global.Question{
			{
				Got:  "babad",
				Want: "bab",
			},
			{
				Got:  "cbbd",
				Want: "bb",
			},
			{
				Got:  "abbcccddcccbba",
				Want: "abbcccddcccbba",
			},
			{
				Got:  "a",
				Want: "a",
			},
		}

		for _, v := range qs {
			assert.Equal(t, strings.LongestPalindrome(v.Got), v.Want)
		}
	})
}
