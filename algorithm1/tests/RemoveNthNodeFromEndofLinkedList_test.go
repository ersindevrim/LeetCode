package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/ersindevrim/leetcode/global"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	t.Run("This should return correct value", func(t *testing.T) {
		postData1 := global.ListNode{
			Val: 1,
			Next: &global.ListNode{
				Val: 2,
				Next: &global.ListNode{
					Val: 3,
					Next: &global.ListNode{
						Val: 4,
						Next: &global.ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		}

		result := global.ListNode{
			Val: 1,
			Next: &global.ListNode{
				Val: 2,
				Next: &global.ListNode{
					Val: 3,
					Next: &global.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		}

		got := algorithm1.RemoveNthFromEnd(&postData1, 2)
		want := &result

		assert.Equal(t, got, want)
	})
}
