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

// InsertToLinkedList ...
func InsertToLinkedList(val int, linked *ListNode) *ListNode {
	newNode := ListNode{
		Val:  val,
		Next: nil,
	}

	if linked == nil {
		linked = &newNode
		return linked
	}

	returnVal := linked
	for returnVal.Next != nil {
		returnVal = returnVal.Next
	}
	returnVal.Next = &newNode
	return linked
}
