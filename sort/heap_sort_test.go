package algo

import "fmt"

func ExampleHeapSort() {
	a := []int{1}
	fmt.Println(HeapSort(a))

	b := []int{5, 2, 3, 8, 1}
	fmt.Println(HeapSort(b))

	c := []int{4, 2, 8, 3}
	fmt.Println(HeapSort(c))

	d := []int{3, 3, 2, 1, 2, 6}
	fmt.Println(HeapSort(d))

	// Output:
	// [1]
	// [1 2 3 5 8]
	// [2 3 4 8]
	// [1 2 2 3 3 6]
}
