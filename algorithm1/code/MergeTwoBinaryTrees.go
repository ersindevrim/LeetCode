package algorithm1

import "github.com/ersindevrim/leetcode/global"

// MergeTrees ...
func MergeTrees(root1 *global.TreeNode, root2 *global.TreeNode) *global.TreeNode {

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	//core is ready
	myRoot := global.TreeNode{Val: root1.Val + root2.Val, Left: nil, Right: nil}
	//left branches
	myRoot.Left = MergeTrees(root1.Left, root2.Left)
	//right branches
	myRoot.Right = MergeTrees(root1.Right, root2.Right)

	return &myRoot
}
