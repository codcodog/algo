package easy

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

func MergeTwoLists(list1, list2 *ListNode) *ListNode {
	prehead := &ListNode{
		Val: -1,
	}

	prev := prehead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}

	if list1 != nil {
		prev.Next = list1
	}
	if list2 != nil {
		prev.Next = list2
	}

	return prehead.Next
}
