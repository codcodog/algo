package data_structure

import (
	"errors"
	"fmt"
)

type Item string

type Node struct {
	data Item
	next *Node
}

type ItemLinkedList struct {
	head *Node
	size int
}

func (ll *ItemLinkedList) Append(item Item) {
	node := Node{item, nil}

	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	ll.size++
}

func (ll *ItemLinkedList) Insert(i int, item Item) error {
	if i < 0 || i > ll.size {
		return fmt.Errorf("Index out of bounds")
	}

	newNode := Node{item, nil}
	if i == 0 {
		newNode.next = ll.head
		ll.head = &newNode
	} else {
		node := ll.head
		for j := 0; j < i-2; j++ {
			node = node.next
		}
		newNode.next = node.next
		node.next = &newNode
	}

	ll.size++
	return nil
}

func (ll *ItemLinkedList) RemoveAt(i int) error {
	if i < 0 || i > ll.size {
		return fmt.Errorf("Index out of bounds")
	}

	preNode := ll.head
	for j := 0; j < i-2; j++ {
		preNode = preNode.next
	}
	deleteNode := preNode.next
	preNode.next = deleteNode.next

	ll.size--
	return nil
}

func (ll *ItemLinkedList) IndexOf(item Item) (int, error) {
	node := ll.head
	for i := 0; i < ll.size; i++ {
		if node.data == item {
			return i, nil
		}
		node = node.next
	}
	err := errors.New("not found")
	return -1, err
}

func (ll *ItemLinkedList) Size() int {
	return ll.size
}

func (ll *ItemLinkedList) IsEmpty() bool {
	if ll.head == nil {
		return true
	}
	return false
}
