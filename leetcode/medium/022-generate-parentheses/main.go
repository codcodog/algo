package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

var ans []string

func generateParenthesis(n int) []string {
	ans = make([]string, 0)
	backtrack("", n, 0, 0)
	return ans
}

func backtrack(s string, n, left, right int) {
	if len(s) == 2*n {
		ans = append(ans, s)
		return
	}
	if left < n {
		backtrack(s+"(", n, left+1, right)
	}
	if right < left {
		backtrack(s+")", n, left, right+1)
	}
}
