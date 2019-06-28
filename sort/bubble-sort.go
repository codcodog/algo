package main

import "fmt"

func main() {
	input := []int{1, 9, 2, 8, 3, 2}
	bubbleSort(input)
	fmt.Println(input)
}

func bubbleSort(list []int) {
	if len(list) < 2 {
		return
	}

	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
