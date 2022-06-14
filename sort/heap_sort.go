package algo

// 堆排序
func HeapSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	dataLen := len(data)
	buildMaxHeap(data, len(data))
	for i := dataLen - 1; i >= 0; i-- {
		data[0], data[i] = data[i], data[0]
		dataLen -= 1
		heapify(data, 0, dataLen)
	}

	return data
}

// 创建最大堆
func buildMaxHeap(data []int, dataLen int) {
	for i := dataLen / 2; i >= 0; i-- {
		heapify(data, i, dataLen)
	}
}

// 堆调整
func heapify(data []int, i, dataLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < dataLen && data[left] > data[largest] {
		largest = left
	}
	if right < dataLen && data[right] > data[largest] {
		largest = right
	}

	if largest != i {
		data[largest], data[i] = data[i], data[largest]
		heapify(data, largest, dataLen)
	}
}
