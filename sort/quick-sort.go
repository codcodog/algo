package main

import "fmt"

func main() {
	input := []int{1, 5, 2, 9, 12, 3, 2}
	quickSort(input)
	fmt.Println(input)
}

// 原地分割 - 快排
func quickSort(list []int) {
	if len(list) < 2 {
		return
	}

	mid := list[0]
	head, tail := 0, len(list)-1
	for i := 1; i <= tail; {
		if list[i] > mid {
			list[i], list[tail] = list[tail], list[i]
			tail--
		} else {
			list[i], list[head] = list[head], list[i]
			head++
			i++
		}
	}
	quickSort(list[:head])
	quickSort(list[head+1:])
}
