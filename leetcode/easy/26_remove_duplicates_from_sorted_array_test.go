package easy

import "testing"

type DuplicatesData struct {
	nums     []int
	expected int
}

func TestRemoveDuplicates(t *testing.T) {
	data := []DuplicatesData{
		DuplicatesData{
			nums:     []int{1, 1, 2},
			expected: 2,
		},
		DuplicatesData{
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, item := range data {
		if RemoveDuplicates(item.nums) != item.expected {
			t.Errorf("test remove duplicates failed: %+v", item)
		}
	}
}
