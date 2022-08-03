package global

// Question ..
type Question struct {
	Got      string
	Got2     string
	IntParam int
	Want     string
	WantBool bool
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsInGrid ...
func IsInGrid(grid [][]int, x, y int) bool {
	return x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0
}
