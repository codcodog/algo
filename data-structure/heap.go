package data_structure

import (
	"errors"
	"fmt"
)

type Heap struct {
	data  []int
	count int
}

func NewHeap() *Heap {
	return &Heap{
		data:  make([]int, 0),
		count: 0}
}

func (h *Heap) Insert(num int) {
	i := h.count
	h.data = append(h.data, num)
	parent := (i - 1) / 2
	for parent >= 0 && h.data[parent] < h.data[i] {
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
		parent = (i - 1) / 2
	}
	h.count++
}

func (h *Heap) Get() int {
	if h.count == 0 {
		return -1
	}
	return h.data[0]
}

func (h *Heap) Delete() error {
	if h.count == 0 {
		return errors.New("empty heap")
	}

	h.data[0], h.data[h.count-1] = h.data[h.count-1], h.data[0]
	h.count--
	h.Heapify()
	return nil
}

func (h *Heap) Heapify() {
	i := 0
	parent := i
	leftChild := 2*i + 1
	rightChild := 2*i + 2
	for {
		if h.data[leftChild] > h.data[parent] {
			h.data[parent], h.data[leftChild] = h.data[leftChild], h.data[parent]
			i = leftChild
		}
		if h.data[rightChild] > h.data[parent] {
			h.data[parent], h.data[rightChild] = h.data[rightChild], h.data[parent]
			i = rightChild
		}
		parent = i
		leftChild = 2*i + 1
		rightChild = 2*i + 2

		if leftChild >= h.count {
			break
		}
	}
}

func (h *Heap) Print() {
	fmt.Println(h.data)
}
