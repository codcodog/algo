package main

import "fmt"

func main() {
	input := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(input)) // expected: 5 => [6, 4, 8, 10, 9]
}

func findUnsortedSubarray(nums []int) int {
	begin := -1
	end := -2
	leftMax := nums[0]
	n := len(nums)
	rightMin := nums[n-1]
	for i := 1; i < n; i++ {
		leftMax = max(nums[i], leftMax)
		rightMin = min(nums[n-1-i], rightMin)
		if nums[i] < leftMax {
			end = i
		}
		if nums[n-1-i] > rightMin {
			begin = n - 1 - i
		}
	}
	return end - begin + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
