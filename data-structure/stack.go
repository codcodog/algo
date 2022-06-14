package data_structure

import "fmt"

type Stack struct {
	data []interface{}
	top  int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0),
		top:  -1}
}

func (s *Stack) Push(v interface{}) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top++
	}

	if s.top > len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	v := s.data[s.top]
	s.top--
	return v
}

func (s *Stack) Print() {
	if s.IsEmpty() {
		fmt.Println("empty stack")
	}

	for i := s.top; i >= 0; i-- {
		fmt.Printf("%d ", s.data[i])
	}
	fmt.Println()
}

func (s *Stack) IsEmpty() bool {
	return s.top < 0
}
