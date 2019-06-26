package main

import (
	"fmt"
)

func main() {
	input := []int{1, 2, 3}
	fmt.Println(permute(input))
}

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	backtrack(nums, []int{})
	return res
}

func backtrack(nums []int, path []int) {
	if len(path) == len(nums) {
		res = append(res, path)
		return
	}

	for index, _ := range nums {
		lst := make([]int, len(path))
		copy(lst, path)

		if contains(lst, nums[index]) {
			continue
		}
		lst = append(lst, nums[index])
		backtrack(nums, lst)
	}
}

func contains(lst []int, num int) bool {
	for _, value := range lst {
		if num == value {
			return true
		}
	}
	return false
}
