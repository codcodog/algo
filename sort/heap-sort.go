package main

import "fmt"

func main() {
	input := []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
	heapSort(input)
	fmt.Println(input)
}

func heapSort(list []int) {
	length := len(list)

	// 从最后一个父结点开始调整堆
	for i := length / 2; i > -1; i-- {
		heap(list, i, length-1)
	}

	// 开始堆排序
	// 把根结点和最后一个结点交换，并把交换后的最后一个结点移出堆
	for i := length - 1; i > 0; i-- {
		list[i], list[0] = list[0], list[i]
		heap(list, 0, i-1)
	}
}

// 最大堆调整
func heap(list []int, start, end int) {
	left := 2*start + 1
	if left > end {
		return
	}

	right := 2*start + 2
	large := left
	if right <= end && list[right] > list[left] {
		large = right
	}
	if list[start] > list[large] {
		return
	}
	list[large], list[start] = list[start], list[large]
	heap(list, large, end)
}
