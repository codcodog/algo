package medium

func LengthOfLongestSubstring(s string) int {
	n, j, maxLength := len(s), 0, 0
	cache := make(map[byte]int)
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(cache, s[i-1])
		}

		for j < n && cache[s[j]] == 0 {
			cache[s[j]]++
			j++
		}

		length := j - i
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
