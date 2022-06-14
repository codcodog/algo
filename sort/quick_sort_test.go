package algo

import "fmt"

func ExampleQuickSort() {
	a := []int{}
	fmt.Println(QuickSort(a))

	b := []int{1}
	fmt.Println(QuickSort(b))

	c := []int{5, 3, 8, 2}
	fmt.Println(QuickSort(c))

	d := []int{9, 3, 8, 5, 3, 2, 1}
	fmt.Println(QuickSort(d))

	// Output:
	// []
	// [1]
	// [2 3 5 8]
	// [1 2 3 3 5 8 9]
}

func ExampleQuickSortInPlace() {
	a := []int{}
	fmt.Println(QuickSortInPlace(a))

	b := []int{1}
	fmt.Println(QuickSortInPlace(b))

	c := []int{5, 3, 8, 2}
	fmt.Println(QuickSortInPlace(c))

	d := []int{9, 3, 8, 5, 3, 2, 1}
	fmt.Println(QuickSortInPlace(d))

	// Output:
	// []
	// [1]
	// [2 3 5 8]
	// [1 2 3 3 5 8 9]
}
