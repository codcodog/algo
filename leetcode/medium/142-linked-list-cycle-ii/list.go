package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cached := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := cached[head]; ok {
			return head
		}
		cached[head] = true
		head = head.Next
	}
	return nil
}
