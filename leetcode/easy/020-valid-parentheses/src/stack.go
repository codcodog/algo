package stack

import (
	"errors"
)

type Stack struct {
	data []rune
	top  int
}

func New() *Stack {
	return &Stack{
		data: make([]rune, 0),
		top:  -1}
}

func (s *Stack) Push(c rune) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top++
	}
	if s.top > len(s.data)-1 {
		s.data = append(s.data, c)
	} else {
		s.data[s.top] = c
	}
}

func (s *Stack) Pop() (rune, error) {
	if s.top < 0 {
		return 0, errors.New("empty stack")
	}
	c := s.data[s.top]
	s.top--
	return c, nil
}

func (s *Stack) IsEmpty() bool {
	return s.top < 0
}
