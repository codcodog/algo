package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB
	for a != nil && b != nil && a != b {
		a = a.Next
		b = b.Next

		if a == b {
			return a
		}

		if a == nil {
			a = headB
		}
		if b == nil {
			b = headA
		}
	}
	return a
}

func getIntersectionNode1(headA, headA *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	cache := make(map[*ListNode]bool)
	a := headA
	for a != nil {
		cache[a] = true
		a = a.Next
	}

	b := headB
	for b != nil {
		if _, ok := cache[b]; ok {
			return b
		}
		b = b.Next
	}
	return nil
}
