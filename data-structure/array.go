package data_structure

import (
	"errors"
	"fmt"
)

var IndexError error
var FullError error

func init() {
	IndexError = errors.New("out of new index range")
	FullError = errors.New("full array")
}

type Array struct {
	data   []int
	length int
}

func NewArray(capacity int) *Array {
	if capacity <= 0 {
		return nil
	}

	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0}
}

func (a *Array) Find(index int) (int, error) {
	if a.IsIndexOutOfRange(index) {
		return -1, IndexError
	}
	return a.data[index], nil
}

func (a *Array) Insert(index, value int) error {
	if a.Len() == cap(a.data) {
		return FullError
	}
	if a.IsIndexOutOfRange(index) {
		return IndexError
	}

	for i := a.length; i < index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = value
	a.length++
	return nil
}

func (a *Array) Delete(index int) error {
	if a.IsIndexOutOfRange(index) {
		return IndexError
	}

	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return nil
}

func (a *Array) IsIndexOutOfRange(index int) bool {
	return index >= cap(a.data)
}

func (a *Array) Len() int {
	return a.length
}

func (a *Array) Print() {
	tpl := ""
	for i := 0; i < a.Len(); i++ {
		tpl += fmt.Sprintf("%d -> %v ", i, a.data[i])
	}
	fmt.Println(tpl)
}
