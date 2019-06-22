package main

import "fmt"

func main() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(input, 0)) // expected: 4
}

// 二分查找法
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) >> 1 // => (lo + hi) / 2
		if target == nums[mid] {
			return mid
		}

		if nums[lo] <= nums[mid] {
			if target < nums[mid] && target >= nums[lo] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
