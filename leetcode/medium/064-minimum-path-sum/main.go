package main

import "fmt"

func main() {
	input := [][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}
	fmt.Println(minPathSum(input)) // expected: 7 => 1->3->1->1->1
}

// DP: S[i][j] = min(S[i-1][j], [S[i][j-1]])
// 需要注意的是：发生在最上面的行 S[i-1][j] 不存在，最左边的列 S[i][j-1] 不存在
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		for j := 0; j < n; j++ {
			tmp[j] = grid[0][0]
		}
		sum[i] = tmp
	}

	for i := 1; i < m; i++ {
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		sum[0][j] = sum[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sum[i][j] = min(sum[i-1][j], sum[i][j-1]) + grid[i][j]
		}
	}

	return sum[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
