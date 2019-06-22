package main

import "fmt"

func main() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search1(input, 0)) // expected: 4
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

// 先找到旋转的位置，再拆分两部分分别求 target
func search1(nums []int, target int) int {
	rotateIndex := findRotateIndex(nums)
	if target == nums[rotateIndex] {
		return rotateIndex
	} else if target < nums[0] {
		index := findTarget(nums[rotateIndex+1:], target)
		if index != -1 {
			index += rotateIndex + 1
		}
		return index
	} else {
		return findTarget(nums[0:rotateIndex], target)
	}
}

func findRotateIndex(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		pivot := (left + right) >> 1
		if nums[pivot] > nums[pivot+1] {
			return pivot
		} else {
			if nums[pivot] < nums[left] {
				right = pivot - 1
			} else {
				left = pivot + 1
			}
		}
	}
	return 0
}

func findTarget(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
