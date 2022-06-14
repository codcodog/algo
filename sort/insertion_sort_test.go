package algo

import "fmt"

func ExampleInsertionSort() {
	a := []int{}
	sortedA := InsertionSort(a)
	fmt.Println(sortedA)

	b := []int{2}
	sortedB := InsertionSort(b)
	fmt.Println(sortedB)

	c := []int{8, 7, 2, 3, 5, 1}
	sortedC := InsertionSort(c)
	fmt.Println(sortedC)

	// Output:
	// []
	// [2]
	// [1 2 3 5 7 8]
}
