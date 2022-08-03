package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/ersindevrim/leetcode/global"
	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	t.Run("This should success ", func(t *testing.T) {
		qs := []global.Question{
			{
				Got:      "aa",
				Got2:     "a",
				WantBool: false,
			},
			{
				Got:      "aa",
				Got2:     "a*",
				WantBool: true,
			},
			{
				Got:      "ab",
				Got2:     ".*",
				WantBool: true,
			},
		}

		for _, v := range qs {
			assert.Equal(t, algorithm1.IsMatch(v.Got, v.Got2), v.WantBool)
		}
	})
}
