package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	s := New()
	s.Print()

	s.Push(1)
	s.Push(3)
	s.Push(5)
	s.Print()
	assert.Equal(t, 3, s.top+1)
}

func TestPop(t *testing.T) {
	s := New()
	none := s.Pop()
	assert.Equal(t, nil, none)

	s.Push(1)
	s.Push(3)
	s.Push(5)

	five := s.Pop()
	assert.Equal(t, 5, five)
	assert.Equal(t, 2, s.top+1)
}
