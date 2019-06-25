package main

import "fmt"

func main() {
	input := []int{5, 3, 8, 2, 7, 1}
	fmt.Println(shellSort(input))
}

// 希尔排序 - 插入排序的一种更高效的改进版本
func shellSort(lst []int) []int {
	if len(lst) <= 1 {
		return lst
	}

	gap := len(lst) >> 1 // => len(lst) / 2
	for gap > 0 {
		for i := gap; i < len(lst); i++ {
			tmp := lst[i]
			j := i
			for j >= gap && lst[j-gap] > tmp {
				lst[j] = lst[j-gap]
				j -= gap
			}
			lst[j] = tmp
		}
		gap = gap >> 1
	}
	return lst
}
