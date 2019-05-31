package main

import (
	"fmt"

	stack "github.com/codcodog/stack/src"
)

func main() {
	chars1 := "()[]{}"
	fmt.Println(isValid(chars1))

	chars2 := "["
	fmt.Println(isValid(chars2))
}

func isValid(chars string) bool {
	var another rune
	var err error

	s := stack.New()
	for _, c := range chars {
		switch c {
		case '(':
			s.Push(')')
		case '[':
			s.Push(']')
		case '{':
			s.Push('}')
		default:
			another, err = s.Pop()
			if err != nil || another != c {
				return false
			}
		}
	}

	return s.IsEmpty()
}
