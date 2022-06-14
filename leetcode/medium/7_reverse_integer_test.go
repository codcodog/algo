package medium

import "testing"

type reverseInt struct {
	n        int
	expected int
}

func TestReverseInt(t *testing.T) {
	data := []reverseInt{
		reverseInt{
			n:        123,
			expected: 321,
		},
		reverseInt{
			n:        -123,
			expected: -321,
		},
		reverseInt{
			n:        120,
			expected: 21,
		},
		reverseInt{
			n:        0,
			expected: 0,
		},
	}

	for _, item := range data {
		if ReverseInt(item.n) != item.expected {
			t.Errorf("test reverse integer failed: %+v", item)
		}
	}
}
