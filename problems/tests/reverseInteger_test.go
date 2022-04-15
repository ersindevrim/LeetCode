package problems

import (
	"testing"

	problems "github.com/ersindevrim/leetcode/problems/code"
	"github.com/stretchr/testify/assert"
)

type Mock struct {
	Got  int
	Want int
}

func TestConvert(t *testing.T) {
	t.Run("This should success ", func(t *testing.T) {
		qs := []Mock{
			{
				Got:  123,
				Want: 321,
			},
			{
				Got:  -125790,
				Want: -97521,
			},
			{
				Got:  1234567899,
				Want: 0,
			},
		}

		for _, v := range qs {
			assert.Equal(t, v.Want, problems.Reverse(v.Got))
		}
	})
}
