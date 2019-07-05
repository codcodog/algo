package main

import "fmt"

func main() {
	input := []int{1, 9, 2, 7, 5, 8}
	fmt.Println(mergeSort(input))
}

func mergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	mid := len(list) / 2
	left := mergeSort(list[0:mid])
	right := mergeSort(list[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
