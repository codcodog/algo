package palindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

// 根据快节点是慢节点的两倍速度找到中间节点
// 然后反转后一半，比较链表前一半是否相等
func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// for odd pointer
	if fast != nil {
		slow = slow.Next
	}

	slow = reverse(slow, nil)
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}

func reverse(cur, prev *ListNode) *ListNode {
	if cur == nil {
		return prev
	}
	next := cur.Next
	cur.Next = prev
	prev = cur
	return reverse(next, prev)
}
