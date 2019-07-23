package main

import "fmt"

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict)) // expected: true
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && contains(wordDict, s[j:i]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func contains(wordDict []string, word string) bool {
	for _, value := range wordDict {
		if value == word {
			return true
		}
	}
	return false
}
