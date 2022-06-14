package algo

// 冒泡排序
func BubbleSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	return data
}
