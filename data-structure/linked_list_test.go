package data_structure

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	ll := ItemLinkedList{}

	assert.Equal(t, 0, ll.Size(), fmt.Sprintf("期望空链表，得到链表大小为：%d", ll.Size()))

	ll.Append(Item("hello"))
	assert.Equal(t, 1, ll.Size(), fmt.Sprintf("链表大小错误，期望 1，得到：%d", ll.Size()))

	ll.Append(Item("world"))
	ll.Append(Item("nihao"))
	ll.Append(Item("shijie"))
	assert.Equal(t, 4, ll.Size(), fmt.Sprintf("链表大小错误，期望 4，得到：%d", ll.Size()))
}

func TestEmpty(t *testing.T) {
	ll := ItemLinkedList{}
	assert.True(t, ll.IsEmpty())

	ll.Append(Item("hello"))
	assert.False(t, ll.IsEmpty())
}

func TestInsert(t *testing.T) {
	ll := ItemLinkedList{}
	ll.Append(Item("hello"))
	ll.Append(Item("world"))

	err := ll.Insert(-1, Item("nihao"))
	assert.EqualError(t, err, "Index out of bounds")

	err = ll.Insert(3, Item("nihao"))
	assert.EqualError(t, err, "Index out of bounds")

	ll.Insert(1, Item("shijie"))
	node := ll.head
	for j := 0; j < 1; j++ {
		node = node.next
	}
	assert.Equal(t, node.data, Item("shijie"))
	assert.Equal(t, 3, ll.size)
}

func TestRemoveAt(t *testing.T) {
	ll := ItemLinkedList{}
	ll.Append(Item("hello"))
	ll.Append(Item("world"))

	err := ll.RemoveAt(-1)
	assert.EqualError(t, err, "Index out of bounds")

	err = ll.RemoveAt(3)
	assert.EqualError(t, err, "Index out of bounds")

	ll.RemoveAt(1)
	assert.Equal(t, 1, ll.size, fmt.Sprintf("删除链表后，大小错误，期望 1，得到：%d", ll.size))
}

func TestIndexOf(t *testing.T) {
	ll := ItemLinkedList{}
	ll.Append(Item("hello"))
	ll.Append(Item("world"))

	index, err := ll.IndexOf(Item("nihao"))
	assert.EqualError(t, err, "not found")

	index, _ = ll.IndexOf(Item("world"))
	assert.Equal(t, 1, index)
}
