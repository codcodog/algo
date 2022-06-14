package medium

func LongestPalindromic(s string) string {
	start, end, n := 0, 0, len(s)
	for i := 0; i < n; i++ {
		left, right := expandAroundCenter(s, i, i)
		if right-left > end-start {
			start, end = left, right
		}

		left, right = expandAroundCenter(s, i, i+1)
		if right-left > end-start {
			start, end = left, right
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
