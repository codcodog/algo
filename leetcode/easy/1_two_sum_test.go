package easy

import (
	"reflect"
	"testing"
)

type TwoSumData struct {
	Nums     []int
	Target   int
	Expected []int
}

func TestTwoSum(t *testing.T) {
	data := [3]TwoSumData{
		TwoSumData{
			Nums:     []int{2, 7, 11, 15},
			Target:   9,
			Expected: []int{0, 1},
		},
		TwoSumData{
			Nums:     []int{3, 2, 4},
			Target:   6,
			Expected: []int{1, 2},
		},
		TwoSumData{
			Nums:     []int{3, 3},
			Target:   6,
			Expected: []int{0, 1},
		},
	}

	for _, item := range data {
		indexes := TwoSum(item.Nums, item.Target)
		if !reflect.DeepEqual(indexes, item.Expected) {
			t.Errorf("two sum test failed: %+v", item)
		}
	}
}
