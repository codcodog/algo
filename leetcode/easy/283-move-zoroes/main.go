package main

import "fmt"

func main() {
	input := []int{0, 1, 0, 3, 12}
	moveZeroes(input)
	fmt.Println(input) // expected: [1,3,12,0,0]
}

func moveZeroes(nums []int) {
	length := len(nums)
	insertPos := 0
	for _, num := range nums {
		if num != 0 {
			nums[insertPos] = num
			insertPos++
		}
	}
	for insertPos < length {
		nums[insertPos] = 0
		insertPos++
	}
}
