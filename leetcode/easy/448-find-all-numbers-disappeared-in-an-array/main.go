package main

import "fmt"

func main() {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(input)) // expected: [5, 6]
}

// 1 <= nums[i] <= n
// n-1 对应 nums 数组的 index
func findDisappearedNumbers(nums []int) []int {
	var val int
	for i := 0; i < len(nums); i++ {
		val = abs(nums[i]) - 1
		if nums[val] > 0 {
			nums[val] = -nums[val]
		}
	}

	disappearedNums := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			disappearedNums = append(disappearedNums, i+1)
		}
	}
	return disappearedNums
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
