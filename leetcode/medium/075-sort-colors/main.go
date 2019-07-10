package main

import "fmt"

func main() {
	input := []int{2, 0, 2, 1, 1, 0}
	sortColors(input)
	fmt.Println(input) // expected: [0, 0, 1, 1, 2, 2]
}

// 荷兰国旗问题
// p0 指针 -> 0 最右边界
// p2 指针 -> 2 最左边界
// cur -> 当前元素序号
func sortColors(nums []int) {
	p0, cur, p2 := 0, 0, len(nums)-1
	for cur <= p2 {
		if nums[cur] == 0 {
			nums[cur], nums[p0] = nums[p0], nums[cur]
			p0++
			cur++
		} else if nums[cur] == 2 {
			nums[cur], nums[p2] = nums[p2], nums[cur]
			p2--
		} else {
			cur++
		}
	}
}
