package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 先找到链表长度，再删除 nth node
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := head
	length := 0
	last := head
	for last != nil {
		length++
		last = last.Next
	}

	if length == n {
		return dummy.Next
	}

	destLength := length - n
	first := head
	for {
		destLength--
		if destLength == 0 {
			break
		}
		first = first.Next
	}
	first.Next = first.Next.Next

	return dummy
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	first := head
	second := head
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return head
}
