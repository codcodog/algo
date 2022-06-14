package easy

import "testing"

type parensData struct {
	s        string
	expected bool
}

func TestIsValidParens(t *testing.T) {
	data := []parensData{
		parensData{
			s:        "()",
			expected: true,
		},
		parensData{
			s:        "()[]{}",
			expected: true,
		},
		parensData{
			s:        "([)]",
			expected: false,
		},
		parensData{
			s:        "{[]}",
			expected: true,
		},
		parensData{
			s:        "[",
			expected: false,
		},
		parensData{
			s:        "((",
			expected: false,
		},
	}

	for _, item := range data {
		if IsValidParens(item.s) != item.expected {
			t.Errorf("test is valid parens failed: %+v", item)
		}
	}
}
