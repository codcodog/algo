package data_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayInsert(t *testing.T) {
	a := NewArray(3)

	err := a.Insert(4, 1)
	assert.EqualError(t, err, "out of new index range")

	a.Insert(0, 1)
	a.Insert(1, 2)
	a.Insert(2, 3)
	assert.Equal(t, 3, a.Len())

	first, _ := a.Find(0)
	second, _ := a.Find(1)
	third, _ := a.Find(2)
	assert.Equal(t, 1, first)
	assert.Equal(t, 2, second)
	assert.Equal(t, 3, third)
}

func TestArrayFind(t *testing.T) {
	a := NewArray(3)

	_, err := a.Find(4)
	assert.EqualError(t, err, "out of new index range")

	a.Insert(0, 3)
	a.Insert(1, 5)
	a.Insert(2, 9)
	value, _ := a.Find(1)
	assert.Equal(t, 5, value)
}

func TestArrayDelete(t *testing.T) {
	a := NewArray(5)
	a.Insert(0, 10)
	a.Insert(1, 11)
	a.Insert(2, 12)
	a.Insert(3, 13)

	a.Delete(1)
	assert.Equal(t, 3, a.Len())

	first, _ := a.Find(0)
	second, _ := a.Find(1)
	third, _ := a.Find(2)
	assert.Equal(t, 10, first)
	assert.Equal(t, 12, second)
	assert.Equal(t, 13, third)
}
