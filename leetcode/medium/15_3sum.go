package medium

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	var ans [][]int
	findCache := make(map[string]struct{})
	n := len(nums)
	for i := 0; i < n; i++ {
		cache := make(map[int]int)
		target := 0 - nums[i]
		for j := i + 1; j < n; j++ {
			if value, ok := cache[nums[j]]; ok {
				if !isFind(nums[i], value, nums[j], findCache) {
					ans = append(ans, []int{nums[i], value, nums[j]})
				}
			}
			cache[target-nums[j]] = nums[j]
		}
	}

	return ans
}

func isFind(a, b, c int, findCache map[string]struct{}) bool {
	data := []int{a, b, c}
	sort.Ints(data)
	key := fmt.Sprintf("%d_%d_%d", data[0], data[1], data[2])
	if _, ok := findCache[key]; ok {
		return true
	}
	findCache[key] = struct{}{}

	return false
}
