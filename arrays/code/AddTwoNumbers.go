package arrays

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	move, head := 0, l1
	for {
		l1.Val += l2.Val + move
		move = int(l1.Val / 10)
		l1.Val = l1.Val % 10

		if l2.Next == nil {
			break
		} else if l1.Next == nil {
			l1.Next = l2.Next
			break
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for move != 0 {
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}
		l1.Next.Val += move

		move = l1.Next.Val / 10
		l1.Next.Val = l1.Next.Val % 10
		l1 = l1.Next
	}

	return head
}
