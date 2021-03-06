package arrays

import (
	"testing"

	arrays "github.com/ersindevrim/leetcode/arrays/code"
	"github.com/ersindevrim/leetcode/global"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("This should return correct value", func(t *testing.T) {
		postData1 := global.ListNode{
			Val: 2,
			Next: &global.ListNode{
				Val: 4,
				Next: &global.ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}

		postData2 := global.ListNode{
			Val: 5,
			Next: &global.ListNode{
				Val: 6,
				Next: &global.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}

		result := global.ListNode{
			Val: 7,
			Next: &global.ListNode{
				Val: 0,
				Next: &global.ListNode{
					Val:  8,
					Next: nil,
				},
			},
		}

		got := arrays.AddTwoNumbers(&postData1, &postData2)
		want := &result

		assert.Equal(t, got, want)
	})

}
