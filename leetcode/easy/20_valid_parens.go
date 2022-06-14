package easy

type stack struct {
	data []byte
	top  int
}

func newStack() *stack {
	return &stack{
		data: make([]byte, 0),
		top:  -1,
	}
}

func (s *stack) push(c byte) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top += 1
	}

	if s.top > len(s.data)-1 {
		s.data = append(s.data, c)
	} else {
		s.data[s.top] = c
	}
}

func (s *stack) pop() (c byte) {
	if s.isEmpty() {
		return c
	}

	c = s.data[s.top]
	s.top -= 1
	return c
}

func (s *stack) isEmpty() bool {
	return s.top < 0
}

func IsValidParens(s string) bool {
	stack := newStack()
	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stack.push(s[i])
		case ')':
			s := stack.pop()
			if s != '(' {
				return false
			}
		case ']':
			s := stack.pop()
			if s != '[' {
				return false
			}
		case '}':
			s := stack.pop()
			if s != '{' {
				return false
			}
		}
	}
	if !stack.isEmpty() {
		return false
	}

	return true
}
