package main

import "fmt"

func main() {
	input := 3
	fmt.Println(numTrees(input)) // expected: 5
}

// F(i,n) = G(i−1) * G(n−i)
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}
