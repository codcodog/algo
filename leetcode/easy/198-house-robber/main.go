package main

import "fmt"

func main() {
	input1 := []int{1, 2, 3, 1}
	fmt.Println(rob(input1)) // expected: 1 + 3 = 4

	input2 := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(input2)) // expected: 2 + 9 + 1 = 12
}

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
