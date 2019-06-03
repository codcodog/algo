package main

import (
	"fmt"

	linkedlist "github.com/codcodog/linkedlist/src"
)

func main() {
	l1 := &linkedlist.Linkedlist{}
	l2 := &linkedlist.Linkedlist{}

	l1.Append(linkedlist.Item(1))
	l1.Append(linkedlist.Item(2))
	l1.Append(linkedlist.Item(4))

	l2.Append(linkedlist.Item(1))
	l2.Append(linkedlist.Item(3))
	l2.Append(linkedlist.Item(5))

	mergeHead := mergeTwoLists(l1.Head, l2.Head)
	for mergeHead != nil {
		fmt.Print(mergeHead.Data, " ")
		mergeHead = mergeHead.Next
	}
}

func mergeTwoLists(l1, l2 *linkedlist.Node) *linkedlist.Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Data < l2.Data {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
