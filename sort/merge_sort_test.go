package algo

import "fmt"

func ExampleMergeSort() {
	a := []int{1}
	fmt.Println(MergeSort(a))

	b := []int{1, 9, 2, 8, 4}
	fmt.Println(MergeSort(b))

	c := []int{8, 2, 3, 5}
	fmt.Println(MergeSort(c))

	d := []int{5, 3, 3, 1, 1, 2}
	fmt.Println(MergeSort(d))

	// Output:
	// [1]
	// [1 2 4 8 9]
	// [2 3 5 8]
	// [1 1 2 3 3 5]
}
