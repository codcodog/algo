package main

import "fmt"

func main() {
	input := []int{3, 2, 3}
	fmt.Println(majorityElement(input)) // expected: 3

	input1 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(input1)) // expected: 2
}

// Boyer-Moore Majority Vote Algorithm
func majorityElement(nums []int) int {
	majority := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			majority = nums[i]
			count = 1
		} else if nums[i] == majority {
			count++
		} else {
			count--
		}
	}
	return majority
}
