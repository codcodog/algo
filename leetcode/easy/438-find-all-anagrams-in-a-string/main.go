package main

import (
	"fmt"
)

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p)) // expected: [0, 6]

	s1 := "abab"
	p1 := "ab"
	fmt.Println(findAnagrams(s1, p1)) // expected: [0, 1, 2]
}

func findAnagrams(s string, p string) []int {
	hash := make(map[byte]int)
	for index, _ := range p {
		if _, ok := hash[p[index]]; ok {
			hash[p[index]]++
		} else {
			hash[p[index]] = 1
		}
	}

	left, right, count := 0, 0, len(p)
	find := make([]int, 0)
	for right < len(s) {
		if _, ok := hash[s[right]]; ok {
			hash[s[right]]--
			if hash[s[right]] >= 0 {
				count--
			}
		}

		if count == 0 {
			find = append(find, left)
		}

		if right == left+len(p)-1 {
			if _, ok := hash[s[left]]; ok {
				if hash[s[left]] >= 0 {
					count++
				}
				hash[s[left]]++
			}
			left++
		}
		right++
	}

	return find
}
