package main

import "fmt"

func main() {
	input := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(input)) // expected: 1
}

var n, m int

func numIslands(grid [][]byte) int {
	n = len(grid)
	if n == 0 {
		return 0
	}
	m = len(grid[0])

	var count int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
