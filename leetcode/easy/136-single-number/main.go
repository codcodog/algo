package main

import "fmt"

func main() {
	input := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(input)) // expected: 4
}

func singleNumber(nums []int) int {
	a := 0
	for _, value := range nums {
		a ^= value
	}
	return a
}
