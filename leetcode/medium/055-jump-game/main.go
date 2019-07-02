package main

import "fmt"

func main() {
	input := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(input))

	input1 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(input1))
}

func canJump(nums []int) bool {
	length := len(nums)
	max := 0
	for i := 0; i <= max; i++ {
		max = getMax(max, i+nums[i])
		if max >= length-1 {
			return true
		}
	}
	return false
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
