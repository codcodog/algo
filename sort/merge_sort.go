package algo

// 归并排序
func MergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	mid := len(data) / 2
	leftData := MergeSort(data[:mid])
	rightData := MergeSort(data[mid:])

	var sortData []int
	i, j := 0, 0
	for i < len(leftData) && j < len(rightData) {
		if leftData[i] < rightData[j] {
			sortData = append(sortData, leftData[i])
			i++
		} else {
			sortData = append(sortData, rightData[j])
			j++
		}
	}
	for ; i < len(leftData); i++ {
		sortData = append(sortData, leftData[i])
	}
	for ; j < len(rightData); j++ {
		sortData = append(sortData, rightData[j])
	}

	return sortData
}
