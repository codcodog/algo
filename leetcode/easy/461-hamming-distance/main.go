package main

import "fmt"

func main() {
	x := 1
	y := 4
	fmt.Println(hammingDistance(x, y)) // expected: 2
}

func hammingDistance(x int, y int) int {
	z := x ^ y
	var d int
	for z > 0 {
		if z%2 == 1 {
			d++
		}
		z = z >> 1
	}
	return d
}
