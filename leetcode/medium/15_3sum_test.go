package medium

import "fmt"

func ExampleThreeSum() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(ThreeSum(nums))

	nums = []int{}
	fmt.Println(ThreeSum(nums))

	nums = []int{0}
	fmt.Println(ThreeSum(nums))

	nums = []int{0, 0, 0}
	fmt.Println(ThreeSum(nums))

	nums = []int{-1, 0, 1}
	fmt.Println(ThreeSum(nums))

	// Output:
	// [[-1 0 1] [-1 2 -1]]
	// []
	// []
	// [[0 0 0]]
	// [[-1 0 1]]
}
