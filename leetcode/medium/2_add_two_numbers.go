package medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	var data []int
	for l != nil {
		data = append(data, l.Val)
		l = l.Next
	}
	fmt.Println(data)
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var list, prev *ListNode

	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		value := (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10

		if prev == nil {
			list = &ListNode{value, nil}
			prev = list
		} else {
			prev.Next = &ListNode{value, nil}
			prev = prev.Next
		}
	}

	if carry > 0 {
		prev.Next = &ListNode{carry, nil}
	}

	return list
}
