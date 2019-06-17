package main

import "fmt"

func main() {
	input := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(input)) // expected: 3

	input1 := "bbbbbbb"
	fmt.Println(lengthOfLongestSubstring(input1)) // expected: 1

	input2 := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(input2)) // expected: 3

	input3 := "b"
	fmt.Println(lengthOfLongestSubstring(input3)) // expected: 1
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	cache := make(map[byte]int)
	ans, left, right := 0, 0, 0
	for ; right < len(s); right++ {
		if _, ok := cache[s[right]]; ok {
			left = max(left, cache[s[right]]+1)
		}
		cache[s[right]] = right
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
