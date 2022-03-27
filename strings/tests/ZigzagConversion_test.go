package strings

import (
	"testing"

	strings "github.com/ersindevrim/leetcode/strings/code"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Run("This should success ", func(t *testing.T) {
		qs := []question{
			{
				got:      "PAYPALISHIRING",
				want:     "PAHNAPLSIIGYIR",
				intParam: 3,
			},
			{
				got:      "PAYPALISHIRING",
				want:     "PINALSIGYAHRPI",
				intParam: 4,
			},
			{
				got:      "A",
				want:     "A",
				intParam: 1,
			},
		}

		for _, v := range qs {
			assert.Equal(t, strings.Convert(v.got, v.intParam), v.want)
		}
	})
}
