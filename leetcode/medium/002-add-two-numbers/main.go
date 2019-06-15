package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	curr := head

	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return head.Next
}
