package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// merge sort
// step1: cut the list to two halves
// stpe2: sort each half
// sort3: merge l1 and l2
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// step1
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	// step2
	l1 := sortList(head)
	l2 := sortList(slow)

	// step3
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	p := &ListNode{}
	l := p
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}

	return l.Next
}
