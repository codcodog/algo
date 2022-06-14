package data_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapInsert(t *testing.T) {
	h := NewHeap()

	h.Insert(7)
	h.Insert(3)
	h.Insert(5)
	h.Insert(1)
	h.Insert(2)
	expected := []int{7, 3, 5, 1, 2}
	assert.Equal(t, expected, h.data)
	assert.Equal(t, h.count, len(h.data))
	assert.Equal(t, 5, h.count)

	h.Insert(8)
	expected = []int{8, 3, 7, 1, 2, 5}
	assert.Equal(t, expected, h.data)
	assert.Equal(t, h.count, len(h.data))
	assert.Equal(t, 6, h.count)

	h.Insert(6)
	expected = []int{8, 3, 7, 1, 2, 5, 6}
	assert.Equal(t, expected, h.data)
	assert.Equal(t, h.count, len(h.data))
	assert.Equal(t, 7, h.count)
}

func TestHeapGet(t *testing.T) {
	h := NewHeap()
	head := h.Get()
	assert.Equal(t, -1, head)

	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	head = h.Get()
	assert.Equal(t, 8, head)
}

func TestHeapDelete(t *testing.T) {
	h := NewHeap()
	err := h.Delete()
	assert.EqualError(t, err, "empty heap")

	h.Insert(7)
	h.Insert(3)
	h.Insert(5)
	h.Insert(1)
	h.Insert(2)
	h.Insert(8)
	expected := []int{8, 3, 7, 1, 2, 5}
	assert.Equal(t, expected, h.data)

	h.Delete()
	expected = []int{7, 3, 5, 1, 2}
	assert.Equal(t, expected, h.data[:h.count])
	assert.Equal(t, 5, h.count)
}
