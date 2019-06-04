package main

import "fmt"

func main() {
	n := climbStairs(6)
	fmt.Println(n)
}

func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	dp := make(map[int]int)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}
