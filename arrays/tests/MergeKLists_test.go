package arrays

import (
	"testing"

	arrays "github.com/ersindevrim/leetcode/arrays/code"
	"github.com/ersindevrim/leetcode/global"
	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	t.Run("This should return correct value", func(t *testing.T) {
		postData := []*global.ListNode{
			{
				Val: 1,
				Next: &global.ListNode{
					Val: 4,
					Next: &global.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
			{
				Val: 1,
				Next: &global.ListNode{
					Val: 3,
					Next: &global.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			{
				Val: 2,
				Next: &global.ListNode{
					Val:  6,
					Next: nil,
				},
			},
		}

		result := global.ListNode{
			Val: 1,
			Next: &global.ListNode{
				Val: 1,
				Next: &global.ListNode{
					Val: 2,
					Next: &global.ListNode{
						Val: 3,
						Next: &global.ListNode{
							Val: 4,
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
				},
			},
		}

		got := arrays.MergeKLists(postData)
		want := &result

		assert.Equal(t, got, want)
	})
}
