package arrays

import (
	"sort"

	"github.com/ersindevrim/leetcode/global"
)

// MergeKLists ...
/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeKLists(lists []*global.ListNode) *global.ListNode {
	arr := make([]int, 0)
	length := len(lists)
	var returnNode *global.ListNode

	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	for _, node := range lists {
		if node != nil {
			for {
				arr = append(arr, node.Val)
				if node.Next == nil {
					break
				}
				node = node.Next
			}
		}
	}
	sort.Ints(arr)

	for _, number := range arr {
		returnNode = global.InsertToLinkedList(number, returnNode)
	}

	return returnNode
}
