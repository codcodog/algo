package medium

import "testing"

type PalindromicSubstring struct {
	s        string
	expected string
}

func TestLongestPalindromic(t *testing.T) {
	data := []PalindromicSubstring{
		PalindromicSubstring{
			s:        "babad",
			expected: "bab",
		},
		PalindromicSubstring{
			s:        "cbbd",
			expected: "bb",
		},
		PalindromicSubstring{
			s:        "bb",
			expected: "bb",
		},
	}

	for _, item := range data {
		if LongestPalindromic(item.s) != item.expected {
			t.Errorf("test longest palindromic failed: %+v", item)
		}
	}
}
