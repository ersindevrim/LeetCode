package algorithm1

import "github.com/ersindevrim/leetcode/global"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// MiddleNode...
func MiddleNode(head *global.ListNode) *global.ListNode {
	linkCount := 0
	positionCount := 0
	linkedList := head
	for {
		linkCount++
		if linkedList.Next == nil {
			break
		}
		linkedList = linkedList.Next
	}

	isEven := linkCount % 2
	middleElement := linkCount / 2
	if isEven == 0 {
		for {
			positionCount++
			if positionCount == middleElement+1 {
				return head
			}
			head = head.Next
		}
	} else {
		for {
			positionCount++
			if positionCount == middleElement+1 {
				return head
			}
			head = head.Next
		}
	}
}
