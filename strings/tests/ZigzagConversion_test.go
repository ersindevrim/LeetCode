package strings

import (
	"testing"

	"github.com/ersindevrim/leetcode/global"
	strings "github.com/ersindevrim/leetcode/strings/code"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Run("This should success ", func(t *testing.T) {
		qs := []global.Question{
			{
				Got:      "PAYPALISHIRING",
				Want:     "PAHNAPLSIIGYIR",
				IntParam: 3,
			},
			{
				Got:      "PAYPALISHIRING",
				Want:     "PINALSIGYAHRPI",
				IntParam: 4,
			},
			{
				Got:      "A",
				Want:     "A",
				IntParam: 1,
			},
		}

		for _, v := range qs {
			assert.Equal(t, strings.Convert(v.Got, v.IntParam), v.Want)
		}
	})
}
