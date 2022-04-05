package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/ersindevrim/leetcode/global"
	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	t.Run("This should success ", func(t *testing.T) {
		qs := []global.Question{
			{
				Got:      "ab",
				Got2:     "abab",
				WantBool: true,
			},
			{
				Got:      "abc",
				Got2:     "cbaebabacd",
				WantBool: true,
			},
			{
				Got:      "ab",
				Got2:     "eidboaoo",
				WantBool: false,
			},
		}

		for _, v := range qs {
			assert.Equal(t, algorithm1.CheckInclusion(v.Got, v.Got2), v.WantBool)
		}
	})
}
