package main

import "fmt"

func main() {
	nums := []int{2, 3, 7, 10}
	found := twoSum(nums, 12)
	fmt.Println(found)
}

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := cache[nums[i]]; ok {
			return []int{cache[nums[i]], i}
		}
		cache[target-nums[i]] = i
	}

	return nil
}
