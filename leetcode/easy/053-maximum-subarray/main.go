package main

import (
	"fmt"
)

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max := maxSubArray(input)
	fmt.Println(max) // expected: 6 -> [4, -1, 2, 1]

	fmt.Println(maxSubArray1(input))
}

// 动态规划解法
func maxSubArray(nums []int) int {
	sum := make(map[int]int)
	sum[0] = nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if sum[i-1] > 0 {
			sum[i] = sum[i-1] + nums[i]
		} else {
			sum[i] = nums[i]
		}
		if sum[i] > max {
			max = sum[i]
		}
	}

	return max
}

func maxSubArray1(nums []int) int {
	max := nums[0]
	maxEnd := nums[0]
	var tmp int

	for i := 1; i < len(nums); i++ {
		tmp = maxEnd + nums[i]
		if tmp > nums[i] {
			maxEnd = tmp
		} else {
			maxEnd = nums[i]
		}

		if max < maxEnd {
			max = maxEnd
		}
	}

	return max
}
