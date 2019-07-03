package main

import "fmt"

func main() {
	input := []int{1, 9, 2, 7, 3, 5, 2}
	selectionSort(input)
	fmt.Println(input)
}

func selectionSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		min := i
		for j := i + 1; j < len(list); j++ {
			if list[min] > list[j] {
				min = j
			}
		}
		if min != i {
			list[min], list[i] = list[i], list[min]
		}
	}
}
