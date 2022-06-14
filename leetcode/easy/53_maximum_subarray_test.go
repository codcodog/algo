package easy

import "testing"

type maxSubArray struct {
	nums     []int
	expected int
}

func TestMaxSubArray(t *testing.T) {
	data := []maxSubArray{
		maxSubArray{
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		maxSubArray{
			nums:     []int{1},
			expected: 1,
		},
		maxSubArray{
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
	}

	for _, item := range data {
		if MaxSubArray(item.nums) != item.expected {
			t.Errorf("test max sub array failed: %+v", item)
		}
	}
}
