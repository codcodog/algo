package main

import "fmt"
import "math"

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	maxProfit := maxProfit(input)
	fmt.Println(maxProfit)
}

func maxProfit(prices []int) int {
	var maxProfit int
	min := math.MaxInt64

	for _, value := range prices {
		if value < min {
			min = value
		}
		maxProfit = max((value - min), maxProfit)
	}
	return maxProfit
}

func max(num1, num2 int) int {
	var max int
	if num1 > num2 {
		max = num1
	} else {
		max = num2
	}
	return max
}
