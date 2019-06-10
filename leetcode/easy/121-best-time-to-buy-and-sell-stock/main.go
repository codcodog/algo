package main

import "fmt"
import "math"

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	maxProfit := maxProfit(input)
	fmt.Println(maxProfit)
}

func maxProfit(prices []int) int {
	min := math.MaxInt64
	var maxProfit int
	var tmp int

	for _, value := range prices {
		if value < min {
			min = value
		}

		tmp = value - min
		if tmp > maxProfit {
			maxProfit = tmp
		}
	}
	return maxProfit
}
