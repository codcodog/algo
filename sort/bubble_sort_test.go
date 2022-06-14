package algo

import "fmt"

func ExampleBubbleSort() {
	a := []int{4, 2, 8, 1, 2, 3, 10, 5, 8}
	sortedA := BubbleSort(a)
	fmt.Println(sortedA)

	b := []int{}
	sortedB := BubbleSort(b)
	fmt.Println(sortedB)

	c := []int{3, 2}
	sortedC := BubbleSort(c)
	fmt.Println(sortedC)

	// Output:
	// [1 2 2 3 4 5 8 8 10]
	// []
	// [2 3]
}
