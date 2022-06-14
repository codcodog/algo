package algo

import "fmt"

func ExampleSelectionSort() {
	a := []int{}
	fmt.Println(SelectionSort(a))

	b := []int{1}
	fmt.Println(SelectionSort(b))

	c := []int{1, 3, 9, 2, 2, 3}
	fmt.Println(SelectionSort(c))

	d := []int{8, 4, 5, 1, 2}
	fmt.Println(SelectionSort(d))

	// Output:
	// []
	// [1]
	// [1 2 2 3 3 9]
	// [1 2 4 5 8]
}
