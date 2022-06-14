package algo

// 选择排序
func SelectionSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	var min int
	for i := 0; i < len(data)-1; i++ {
		min = i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		if min != i {
			data[i], data[min] = data[min], data[i]
		}
	}

	return data
}
