package main

import (
	"fmt"
)

func main() {
	m, n := 3, 2
	fmt.Println(uniquePaths(m, n)) // expected: 3

	m, n = 7, 3
	fmt.Println(uniquePaths(m, n)) // expected: 28
}

// key point
// dp[i][j] = dp[i-1][j] + dp[i][j-1]
// dp[0][j] = 1; dp[i][0] = 1
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	dp[0] = make([]int, n)
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
