package arrays

import (
	"testing"

	arrays "github.com/ersindevrim/leetcode/arrays/code"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("This should return correct value", func(t *testing.T) {
		postData1 := arrays.ListNode{
			Val: 2,
			Next: &arrays.ListNode{
				Val: 4,
				Next: &arrays.ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}

		postData2 := arrays.ListNode{
			Val: 5,
			Next: &arrays.ListNode{
				Val: 6,
				Next: &arrays.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}

		result := arrays.ListNode{
			Val: 7,
			Next: &arrays.ListNode{
				Val: 0,
				Next: &arrays.ListNode{
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
