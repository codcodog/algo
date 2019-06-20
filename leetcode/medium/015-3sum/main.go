package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(input)) // expected: [[-1, 0, 1], [-1, -1, 2]]
}

// 先排序，再用双指针
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			curSum := nums[i] + nums[left] + nums[right]
			if curSum == 0 {
				tmp := []int{nums[i], nums[left], nums[right]}
				result = append(result, tmp)

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if curSum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return result
}
