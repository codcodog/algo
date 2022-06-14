package algo

// 插入排序
func InsertionSort(data []int) []int {
	num := len(data)
	if num < 2 {
		return data
	}

	for i := 1; i < num; i++ {
		value := data[i]
		j := i - 1
		for j >= 0 && data[j] > value {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = value
	}

	return data
}
