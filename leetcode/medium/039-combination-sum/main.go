package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(input, target)) // expected: [[7], [2, 2, 3]]
}

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0)
	backtrack(candidates, target, 0, path)
	return res
}

func backtrack(nums []int, target int, index int, path []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		res = addSet(res, path)
		return
	}

	for i := range nums {
		newPath := make([]int, len(path))
		copy(newPath, path)
		backtrack(nums, target-nums[i], i, append(newPath, nums[i]))
	}
}

func addSet(dest [][]int, member []int) [][]int {
	sort.Ints(member)
	for _, value := range dest {
		if equal(value, member) {
			return dest
		}
	}
	return append(dest, member)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for index, value := range a {
		if value != b[index] {
			return false
		}
	}
	return true
}
