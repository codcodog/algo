package main

import "fmt"

func main() {
	input := "babad"
	fmt.Println(longestPalindrome(input)) // expected: bab

	input1 := "cbbd"
	fmt.Println(longestPalindrome(input1)) // expected: bb
}

var lo, maxLen int

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	lo, maxLen = 0, 0
	for i := 0; i < len(s)-1; i++ {
		expandAroundCenter(s, i, i)   // for odd length
		expandAroundCenter(s, i, i+1) // for even length
	}
	return s[lo : lo+maxLen]
}

func expandAroundCenter(s string, left, right int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	if maxLen < right-left-1 {
		lo = left + 1
		maxLen = right - left - 1
	}
}
