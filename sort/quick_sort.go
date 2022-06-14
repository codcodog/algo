package algo

import "math/rand"

// 快速排序
func QuickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	var left, right []int
	base := data[0]
	for index, value := range data {
		if index == 0 { // 跳过，0 下标的数值为基数了，不进行对比
			continue
		}
		if value <= base {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	var resp []int
	resp = append(resp, QuickSort(left)...)
	resp = append(resp, base)
	resp = append(resp, QuickSort(right)...)

	return resp
}

// 快速排序，原地分割版本
func QuickSortInPlace(data []int) []int {
	if len(data) < 2 {
		return data
	}

	left, right := 0, len(data)-1
	pivot := rand.Int() % len(data)

	data[right], data[pivot] = data[pivot], data[right]

	for i, _ := range data {
		if data[i] < data[right] {
			data[left], data[i] = data[i], data[left]
			left++
		}
	}
	data[left], data[right] = data[right], data[left]

	QuickSortInPlace(data[:left])
	QuickSortInPlace(data[left+1:])

	return data
}
