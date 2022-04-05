package algorithm1

import (
	"testing"

	algorithm1 "github.com/ersindevrim/leetcode/algorithm1/code"
	"github.com/ersindevrim/leetcode/global"
	"github.com/stretchr/testify/assert"
)

func TestMergeTrees(t *testing.T) {
	t.Run("This should return correct tree", func(t *testing.T) {
		root1 := global.TreeNode{
			Val: 1,
			Left: &global.TreeNode{
				Val: 3,
				Left: &global.TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &global.TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}

		root2 := global.TreeNode{
			Val: 2,
			Left: &global.TreeNode{
				Val:  1,
				Left: nil,
				Right: &global.TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &global.TreeNode{
				Val:  3,
				Left: nil,
				Right: &global.TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}

		mockWant := &global.TreeNode{
			Val: 3,
			Left: &global.TreeNode{
				Val: 4,
				Left: &global.TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &global.TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &global.TreeNode{
				Val:  5,
				Left: nil,
				Right: &global.TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}

		got := algorithm1.MergeTrees(&root1, &root2)

		//assert.Equal(t, got, mockWant)
		assert.EqualValues(t, got, mockWant)
	})
}
