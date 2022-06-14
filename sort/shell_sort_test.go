package algo

import "fmt"

func ExampleShellSort() {
	a := []int{1}
	fmt.Println(ShellSort(a))

	b := []int{4, 2, 8, 1, 3}
	fmt.Println(ShellSort(b))

	c := []int{9, 3, 2, 7}
	fmt.Println(ShellSort(c))

	d := []int{4, 2, 2, 3, 3, 1}
	fmt.Println(ShellSort(d))

	// Output:
	// [1]
	// [1 2 3 4 8]
	// [2 3 7 9]
	// [1 2 2 3 3 4]
}
