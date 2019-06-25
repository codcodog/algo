package main

import "fmt"

func main() {
	input := []int{5, 3, 8, 2, 7, 1}
	fmt.Println(insertSort(input))
}

// 插入排序
func insertSort(lst []int) []int {
	if len(lst) <= 1 {
		return lst
	}

	for i := 1; i < len(lst); i++ {
		tmp := lst[i]
		j := i - 1

		for j >= 0 && tmp < lst[j] {
			lst[j+1] = lst[j]
			j--
		}
		lst[j+1] = tmp
	}

	return lst
}
