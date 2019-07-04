package main

import "fmt"
import "sort"

func main() {
	input := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	fmt.Println(merge(input))

	input1 := [][]int{[]int{1, 4}, []int{4, 5}}
	fmt.Println(merge(input1))
}

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		n := len(result)
		if n > 0 && result[n-1][1] >= intervals[i][0] {
			result[n-1][1] = max(intervals[i][1], result[n-1][1])
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
