package medium

import "testing"

type LongestSubstring struct {
	s        string
	expected int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	data := []LongestSubstring{
		LongestSubstring{
			s:        "abcaabcbb",
			expected: 3,
		},
		LongestSubstring{
			s:        "bbbbb",
			expected: 1,
		},
		LongestSubstring{
			s:        "pwwkew",
			expected: 3,
		},
		LongestSubstring{
			s:        " ",
			expected: 1,
		},
	}

	for _, item := range data {
		if LengthOfLongestSubstring(item.s) != item.expected {
			t.Errorf("test longest substring failed: %+v", item)
		}
	}
}
