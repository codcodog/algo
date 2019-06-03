package linkedlist

import "fmt"

type Item int

type Node struct {
	Data Item
	Next *Node
}

type Linkedlist struct {
	Head *Node
}

func (l *Linkedlist) Append(data Item) {
	node := &Node{data, nil}

	if l.Head == nil {
		l.Head = node
	} else {
		last := l.Head
		for last.Next != nil {
			last = last.Next
		}
		last.Next = node
	}
}

func (l *Linkedlist) Print() {
	last := l.Head
	for last != nil {
		fmt.Printf("%d ", last.Data)
		last = last.Next
	}
}
