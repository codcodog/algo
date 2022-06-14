package easy

import (
	"testing"
)

type romanData struct {
	s        string
	expected int
}

func TestRomanToInt(t *testing.T) {
	data := []romanData{
		romanData{
			s:        "III",
			expected: 3,
		},
		romanData{
			s:        "IV",
			expected: 4,
		},
		romanData{
			s:        "IX",
			expected: 9,
		},
		romanData{
			s:        "LVIII",
			expected: 58,
		},
		romanData{
			s:        "MCMXCIV",
			expected: 1994,
		},
	}

	for _, item := range data {
		if RomanToInt(item.s) != item.expected {
			t.Errorf("test roman to int failed: %+v", item)
		}
		if RomanToIntV1(item.s) != item.expected {
			t.Errorf("test roman to int failed: %+v", item)
		}
	}
}
