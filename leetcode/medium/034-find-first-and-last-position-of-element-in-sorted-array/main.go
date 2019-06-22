package main

import "fmt"

func main() {
	input := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(input, 8)) // expected: [3, 4]
}

func searchRange(nums []int, target int) []int {
	leftIndex := extremeInsertionIndex(nums, target, true)

	if leftIndex == len(nums) || nums[leftIndex] != target {
		return []int{-1, -1}
	}
	return []int{leftIndex, extremeInsertionIndex(nums, target, false) - 1}
}

func extremeInsertionIndex(nums []int, target int, left bool) int {
	lo, hi := 0, len(nums)

	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] > target || (left && target == nums[mid]) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
