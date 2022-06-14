package algo

// 希尔排序
func ShellSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	gap := len(data) / 2
	for gap > 0 {
		for i := gap; i < len(data); i++ {
			j := i
			for j >= gap && data[j] < data[j-gap] {
				data[j], data[j-gap] = data[j-gap], data[j]
				j = j - gap
			}
		}
		gap = gap / 2
	}

	return data
}
