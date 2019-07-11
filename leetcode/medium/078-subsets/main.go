package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{1, 2, 3}
	fmt.Println(subsets(input))
}

var result [][]int

func subsets(nums []int) [][]int {
	sort.Ints(nums)

	result = make([][]int, 0)
	tmplist := make([]int, 0)
	start := 0
	backtrack(tmplist, nums, start)
	return result
}

func backtrack(list []int, nums []int, start int) {
	result = append(result, list)
	tmp := make([]int, len(list))
	copy(tmp, list)
	for i := start; i < len(nums); i++ {
		backtrack(append(tmp, nums[i]), nums, i+1)
	}
}
