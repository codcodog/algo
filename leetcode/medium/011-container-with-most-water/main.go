package main

import "fmt"

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(input)) // expected: 49
}

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1
	for left != right {
		area := min(height[left], height[right]) * (right - left)
		maxArea = max(area, maxArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
