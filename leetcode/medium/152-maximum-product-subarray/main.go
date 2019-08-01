package main

import "fmt"
import "math"

func main() {
	input := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(input)) // expected: 6 => [2, 3]

	input1 := []int{-2, 0, -1}
	fmt.Println(maxProduct(input1)) // expected: 0
}

func maxProduct(nums []int) int {
	max, imax, imin := math.MinInt64, 1, 1

	for _, value := range nums {
		if value < 0 {
			imax, imin = imin, imax
		}
		imax = Max(imax*value, value)
		imin = min(imin*value, value)
		max = Max(max, imax)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
