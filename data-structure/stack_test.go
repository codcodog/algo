package data_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	s := NewStack()
	s.Print()

	s.Push(1)
	s.Push(3)
	s.Push(5)
	s.Print()
	assert.Equal(t, 3, s.top+1)
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	none := s.Pop()
	assert.Equal(t, nil, none)

	s.Push(1)
	s.Push(3)
	s.Push(5)

	five := s.Pop()
	assert.Equal(t, 5, five)
	assert.Equal(t, 2, s.top+1)
}
