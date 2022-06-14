package easy

import "testing"

type stairs struct {
	n        int
	expected int
}

func TestClimbStairs(t *testing.T) {
	data := []stairs{
		stairs{
			n:        2,
			expected: 2,
		},
		stairs{
			n:        3,
			expected: 3,
		},
	}

	for _, item := range data {
		if ClimbStairs(item.n) != item.expected {
			t.Errorf("test climbing stairs failed: %+v", item)
		}
		if ClimbStairsV1(item.n) != item.expected {
			t.Errorf("test climbing stairs failed: %+v", item)
		}
	}
}
