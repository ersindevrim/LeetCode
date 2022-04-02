package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/ersindevrim/leetcode/global"
	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
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
			Val: 3,
			Next: &global.ListNode{
				Val: 4,
				Next: &global.ListNode{
					Val:  5,
					Next: nil,
				},
			},
		}

		got := algorithm1.MiddleNode(&postData1)
		want := &result

		assert.Equal(t, got, want)
	})

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
							Val: 5,
							Next: &global.ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		}

		result := global.ListNode{
			Val: 4,
			Next: &global.ListNode{
				Val: 5,
				Next: &global.ListNode{
					Val:  6,
					Next: nil,
				},
			},
		}

		got := algorithm1.MiddleNode(&postData1)
		want := &result

		assert.Equal(t, got, want)
	})
}
